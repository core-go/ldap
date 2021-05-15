package ldap

import (
	"context"
	"crypto/tls"
	"fmt"
	"gopkg.in/ldap.v3"
	"strconv"
	"time"
)

type LDAPInfoLoader interface {
	GetLdapInfo(ctx context.Context, id string) (map[string]interface{}, error)
}
type LDAPService struct {
	Config LDAPConfig
}

func NewLDAPService(conf LDAPConfig) *LDAPService {
	if len(conf.Filter) == 0 {
		conf.Filter = "userPrincipalName"
	}
	return &LDAPService{Config: conf}
}
func NewConn(c LDAPConfig) (*ldap.Conn, error) {
	var l *ldap.Conn
	var err error
	if c.Timeout > 0 {
		ldap.DefaultTimeout = time.Duration(c.Timeout) * time.Millisecond
	}
	if c.TLS {
		if c.InsecureSkipVerify {
			l, err = ldap.DialTLS("tcp", c.Server, &tls.Config{ServerName: c.Server, InsecureSkipVerify: true})
		} else {
			l, err = ldap.DialTLS("tcp", c.Server, &tls.Config{ServerName: c.Server})
		}
	} else {
		l, err = ldap.Dial("tcp", c.Server)
		if err == nil {
			if c.StartTLS {
				if c.InsecureSkipVerify {
					err = l.StartTLS(&tls.Config{ServerName: c.Server, InsecureSkipVerify: true})
				} else {
					err = l.StartTLS(&tls.Config{ServerName: c.Server})
				}
			}
		}
	}
	return l, err
}
func (s *LDAPService) GetLdapInfo(ctx context.Context, id string) (map[string]interface{}, error) {
	l, er1 := NewConn(s.Config)
	if er1 != nil {
		return nil, er1
	}
	defer l.Close()
	er2 := l.Bind(s.Config.Username, s.Config.Password)
	if er2 != nil {
		return nil, er2
	}
	searchAttribute := BuildSearchAttributes(s.Config.Attributes, s.Config.Dates)
	searchRequest := ldap.NewSearchRequest(
		s.Config.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 1, 0, false,
		fmt.Sprintf("(&(%s=%s))", s.Config.Filter, id), searchAttribute, nil)
	sr, er3 := l.Search(searchRequest)
	if er3 != nil {
		return nil, er3
	}
	if len(sr.Entries) >= 1 {
		entry := sr.Entries[0]
		return BuildResult(entry, s.Config.Attributes, s.Config.Dates), nil
	} else {
		return nil, nil
	}
}
func BuildSearchAttributes(conf map[string]string, dates map[string]string) []string {
	var searchAttribute []string
	for _, e := range conf {
		searchAttribute = append(searchAttribute, e)
	}
	if dates != nil && len(dates) > 0 {
		for _, d := range conf {
			searchAttribute = append(searchAttribute, d)
		}
	}
	return searchAttribute
}
func BuildResult(entry *ldap.Entry, conf map[string]string, dates map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, e := range conf {
		s := entry.GetAttributeValue(e)
		if len(s) > 0 {
			result[k] = s
		}
	}
	if dates != nil && len(dates) > 0 {
		for k, e := range dates {
			s := entry.GetAttributeValue(e)
			if len(s) > 0 {
				d := ToDate(s)
				if d != nil {
					result[k] = d
				}
			}
		}
	}
	return result
}

const u = 11644473600
const nd = "9223372036854775807"

func ToDate(v string) *time.Time {
	if v == nd {
		return nil
	}
	i, er := strconv.ParseInt(v, 10, 64)
	if er != nil {
		return nil
	}
	l := i / 10000000
	x := time.Unix(l-u, 0)
	return &x
}

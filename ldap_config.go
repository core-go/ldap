package ldap

type LDAPConfig struct {
	Server             string            `mapstructure:"server" json:"server,omitempty" gorm:"column:server" bson:"server,omitempty" dynamodbav:"server,omitempty" firestore:"server,omitempty"`
	BaseDN             string            `mapstructure:"base_dn" json:"baseDN,omitempty" gorm:"column:basedn" bson:"baseDN,omitempty" dynamodbav:"baseDN,omitempty" firestore:"baseDN,omitempty"`
	Timeout            int64             `mapstructure:"timeout" json:"timeout,omitempty" gorm:"column:timeout" bson:"timeout,omitempty" dynamodbav:"timeout,omitempty" firestore:"timeout,omitempty"`
	Domain             string            `mapstructure:"domain" json:"domain,omitempty" gorm:"column:domain" bson:"domain,omitempty" dynamodbav:"domain,omitempty" firestore:"domain,omitempty"`
	Username           string            `mapstructure:"username" json:"username,omitempty" gorm:"column:username" bson:"username,omitempty" dynamodbav:"username,omitempty" firestore:"username,omitempty"`
	Password           string            `mapstructure:"password" json:"password,omitempty" gorm:"column:password" bson:"password,omitempty" dynamodbav:"password,omitempty" firestore:"password,omitempty"`
	Filter             string            `mapstructure:"filter" json:"filter,omitempty" gorm:"column:filter" bson:"filter,omitempty" dynamodbav:"filter,omitempty" firestore:"filter,omitempty"`
	TLS                bool              `mapstructure:"tls" json:"tls,omitempty" gorm:"column:tls" bson:"tls,omitempty" dynamodbav:"tls,omitempty" firestore:"tls,omitempty"`
	StartTLS           bool              `mapstructure:"start_tls" json:"startTLS,omitempty" gorm:"column:starttls" bson:"startTLS,omitempty" dynamodbav:"startTLS,omitempty" firestore:"startTLS,omitempty"`
	InsecureSkipVerify bool              `mapstructure:"insecure_skip_verify" json:"insecureSkipVerify,omitempty" gorm:"column:insecureskipverify" bson:"insecureSkipVerify,omitempty" dynamodbav:"insecureSkipVerify,omitempty" firestore:"insecureSkipVerify,omitempty"`
	Attributes         map[string]string `mapstructure:"attributes" json:"attributes,omitempty" gorm:"column:attributes" bson:"attributes,omitempty" dynamodbav:"attributes,omitempty" firestore:"attributes,omitempty"`
	Dates              map[string]string `mapstructure:"dates" json:"dates,omitempty" gorm:"column:dates" bson:"dates,omitempty" dynamodbav:"dates,omitempty" firestore:"dates,omitempty"`
}

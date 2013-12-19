package mandrill

type Message struct {
	Html                    string               `json:"html,omitempty"`
	Text                    string               `json:"text,omitempty"`
	Subject                 string               `json:"subject"`
	FromEmail               string               `json:"from_email"`
	FromName                string               `json:"from_name"`
	To                      []To                 `json:"to"`
	Headers                 map[string]string    `json:"headers,omitempty"`
	Important               bool                 `json:"important,omitempty"`
	TrackOpens              bool                 `json:"track_opens,omitempty"`
	TrackClicks             bool                 `json:"track_clicks,omitempty"`
	AutoText                bool                 `json:"auto_text,omitempty"`
	AutoHtml                bool                 `json:"auto_html,omitempty"`
	InlineCss               bool                 `json:"inline_css,omitempty"`
	UrlStripQS              bool                 `json:"url_strip_qs,omitempty"`
	PreserveRecipients      bool                 `json:"preserve_recipients,omitempty"`
	ViewContentLink         bool                 `json:"view_content_link,omitempty"`
	BCCAddress              string               `json:"bcc_address,omitempty"`
	TrackingDomain          string               `json:"tracking_domain,omitempty"`
	SigningDomain           string               `json:"signing_domain,omitempty"`
	ReturnPathDomain        string               `json:"return_path_domain,omitempty"`
	Merge                   bool                 `json:"merge,omitempty"`
	GlobalMergeVars         []*GlobalVar         `json:"global_merge_vars,omitempty"`
	MergeVars               []*MergeVars         `json:"merge_vars,omitempty"`
	Tags                    []string             `json:"tags,omitempty"`
	Subaccount              string               `json:"subaccount,omitempty"`
	GoogleAnalyticsDomains  []string             `json:"google_analytics_domains,omitempty"`
	GoogleAnalyticsCampaign []string             `json:"google_analytics_campaign,omitempty"`
	Metadata                []map[string]string  `json:"metadata,omitempty"`
	RecipientMetadata       []*RecipientMetaData `json:"recipient_metadata,omitempty"`
	Attachments             []*Attachment        `json:"attachments,omitempty"`
	Images                  []*Image             `json:"images,omitempty"`
}

type To struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type MergeVars struct {
	Recipient string      `json:"rcpt,omitempty"`
	Vars      []GlobalVar `json:"vars,omitempty"`
}

type GlobalVar struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type TemplateContent struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type Attachment struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type Image struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type RecipientMetaData struct {
	Recipient string            `json:"rcpt,omitempty"`
	Vars      map[string]string `json:"values,omitempty"`
}

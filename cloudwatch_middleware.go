type CloudWatchEvent struct {
	eventVersion 	string 	`json:"eventVersion"`
	UserIdentity
	eventTime 		string 	`json:"eventTime"`
	eventSource		string	`json:"eventSource"`
}

type JSONLogObject struct {
	Date      time.Time         `json:"date"`
	DateDay   string            `json:"date_day"`
	DateMonth string            `json:"date_month"`
	DateTime  string            `json:"date_time"`
	Hostname  string            `json:"hostname"`
	Message   string            `json:"message"`
	PID       string            `json:"pid"`
	Rig       map[string]string `json:"rig"`
	Syslog    map[string]string `json:"syslog"`
	Version   string            `json:"version"`
}

{
  "eventVersion": "1.05",
  "userIdentity": {
    "type": "AssumedRole",
    "principalId": "AROAJD7M6BBLUA5FSXP6Q:rowan.cota",
    "arn": "arn:aws:sts::061305161530:assumed-role/admin/rowan.cota",
    "accountId": "061305161530",
    "accessKeyId": "ASIAQ4RQSQ45PTKNTI3D",
    "sessionContext": {
      "attributes": {
        "mfaAuthenticated": "true",
        "creationDate": "2019-03-27T18:39:05Z"
      },
      "sessionIssuer": {
        "type": "Role",
        "principalId": "AROAJD7M6BBLUA5FSXP6Q",
        "arn": "arn:aws:iam::061305161530:role/admin",
        "accountId": "061305161530",
        "userName": "admin"
      }
    }
  },
  "eventTime": "2019-03-27T19:00:38Z",
  "eventSource": "secretsmanager.amazonaws.com",
  "eventName": "GetSecretValue",
  "awsRegion": "us-east-1",
  "sourceIPAddress": "65.254.2.50",
  "userAgent": "aws-internal/3 aws-sdk-java/1.11.479 Linux/4.9.137-0.1.ac.218.74.329.metal1.x86_64 OpenJDK_64-Bit_Server_VM/25.192-b12 java/1.8.0_192",
  "requestParameters": {
    "secretId": "unstable/rig_controller_api/JWT_signing_key"
  },
  "responseElements": null,
  "requestID": "be111de1-91f7-4977-abe9-908ef87c75da",
  "eventID": "1305fd14-18cd-4b0f-a137-8c3ee53ec6aa",
  "eventType": "AwsApiCall",
  "recipientAccountId": "061305161530"
}
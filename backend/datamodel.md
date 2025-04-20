# User
- id: string
- email: string
- createdDate: timestamp
- registrationType: string (creds/google/..others..) 
- passHash: string
- mobileNumber: string
- emailVerified: bool
- mobileNumberVerified: bool
- userStatus: smallint
  
# Account
Primary account entity
- id: string
- createdDate: timestamp
- primaryUserId

# AccountRole
- userId: string
- accountId: string
- roleId: string

# 

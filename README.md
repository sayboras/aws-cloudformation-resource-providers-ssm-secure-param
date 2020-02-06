# RP::SSM::SecureParameter

# Deployment
1. Login to your AWS account
2. Run `make deploy` to deploy the resource. The newly created resource will be set as default version.
3. Start using `RP::SSM::SecureParameter` in your CF and have fun.

All basic operations `CREATE`, `UPDATE`, `DELETE` are supported.

# Examples

```json
AWSTemplateFormatVersion: 2010-09-09
Description: 'Sample CF for Secure SSM'

Resources:
  SampleSecureSSMParam:
    Type: RP::SSM::SecureParameter
    Properties:
      Description: "Sample secure ssm"
      Name: "/sample/secure/ssm"
      Value: "another value"
      KeyId: "alias/aws/ssm"
      Tags:
        - Key: "key"
          Value: "value"
```
// https://github.com/terraform-aws-modules/terraform-aws-sns/blob/master/examples/complete/main.tf

resource "aws_kms_key" "sns_kms_key" {}

module "sns" {
  source  = "terraform-aws-modules/sns/aws"
  version = "~> 3.0"

  name_prefix       = "users-encrypted-"
  display_name      = "users-encrypted"
  kms_master_key_id = aws_kms_key.sns_kms_key.id
}
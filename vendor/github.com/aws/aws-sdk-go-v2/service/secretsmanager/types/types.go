// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Allows you to filter your list of secrets.
type Filter struct {

	// Filters your list of secrets by a specific key.
	Key FilterNameStringType

	// Filters your list of secrets by a specific value.
	Values []string
}

// A structure that defines the rotation configuration for the secret.
type RotationRulesType struct {

	// Specifies the number of days between automatic scheduled rotations of the
	// secret. Secrets Manager schedules the next rotation when the previous one is
	// complete. Secrets Manager schedules the date by adding the rotation interval
	// (number of days) to the actual date of the last rotation. The service chooses
	// the hour within that 24-hour date window randomly. The minute is also chosen
	// somewhat randomly, but weighted towards the top of the hour and influenced by a
	// variety of factors that help distribute load.
	AutomaticallyAfterDays int64
}

// A structure that contains the details about a secret. It does not include the
// encrypted SecretString and SecretBinary values. To get those values, use the
// GetSecretValue operation.
type SecretListEntry struct {

	// The Amazon Resource Name (ARN) of the secret. For more information about ARNs in
	// Secrets Manager, see Policy Resources
	// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#iam-resources)
	// in the AWS Secrets Manager User Guide.
	ARN *string

	// The date and time when a secret was created.
	CreatedDate *time.Time

	// The date and time the deletion of the secret occurred. Not present on active
	// secrets. The secret can be recovered until the number of days in the recovery
	// window has passed, as specified in the RecoveryWindowInDays parameter of the
	// DeleteSecret operation.
	DeletedDate *time.Time

	// The user-provided description of the secret.
	Description *string

	// The ARN or alias of the AWS KMS customer master key (CMK) used to encrypt the
	// SecretString and SecretBinary fields in each version of the secret. If you don't
	// provide a key, then Secrets Manager defaults to encrypting the secret fields
	// with the default KMS CMK, the key named awssecretsmanager, for this account.
	KmsKeyId *string

	// The last date that this secret was accessed. This value is truncated to midnight
	// of the date and therefore shows only the date, not the time.
	LastAccessedDate *time.Time

	// The last date and time that this secret was modified in any way.
	LastChangedDate *time.Time

	// The last date and time that the rotation process for this secret was invoked.
	LastRotatedDate *time.Time

	// The friendly name of the secret. You can use forward slashes in the name to
	// represent a path hierarchy. For example, /prod/databases/dbserver1 could
	// represent the secret for a server named dbserver1 in the folder databases in the
	// folder prod.
	Name *string

	// Returns the name of the service that created the secret.
	OwningService *string

	// Indicates whether automatic, scheduled rotation is enabled for this secret.
	RotationEnabled bool

	// The ARN of an AWS Lambda function invoked by Secrets Manager to rotate and
	// expire the secret either automatically per the schedule or manually by a call to
	// RotateSecret.
	RotationLambdaARN *string

	// A structure that defines the rotation configuration for the secret.
	RotationRules *RotationRulesType

	// A list of all of the currently assigned SecretVersionStage staging labels and
	// the SecretVersionId attached to each one. Staging labels are used to keep track
	// of the different versions during the rotation process. A version that does not
	// have any SecretVersionStage is considered deprecated and subject to deletion.
	// Such versions are not included in this list.
	SecretVersionsToStages map[string][]string

	// The list of user-defined tags associated with the secret. To add tags to a
	// secret, use TagResource. To remove tags, use UntagResource.
	Tags []Tag
}

// A structure that contains information about one version of a secret.
type SecretVersionsListEntry struct {

	// The date and time this version of the secret was created.
	CreatedDate *time.Time

	// The date that this version of the secret was last accessed. Note that the
	// resolution of this field is at the date level and does not include the time.
	LastAccessedDate *time.Time

	// The unique version identifier of this version of the secret.
	VersionId *string

	// An array of staging labels that are currently associated with this version of
	// the secret.
	VersionStages []string
}

// A structure that contains information about a tag.
type Tag struct {

	// The key identifier, or name, of the tag.
	Key *string

	// The string value associated with the key of the tag.
	Value *string
}

// Displays errors that occurred during validation of the resource policy.
type ValidationErrorsEntry struct {

	// Checks the name of the policy.
	CheckName *string

	// Displays error messages if validation encounters problems during validation of
	// the resource policy.
	ErrorMessage *string
}

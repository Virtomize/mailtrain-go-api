package gomailtrain

const (
	// UnSubscribeOneStep - One-step (i.e. no email with confirmation link)
	UnSubscribeOneStep = 0
	// UnSubscribeOneStepForm - One-step with unsubscription form (i.e. no email with confirmation link)
	UnSubscribeOneStepForm = 1
	// UnSubscribeTwoStep - Two-step (i.e. an email with confirmation link will be sent)
	UnSubscribeTwoStep = 2
	// UnSubscribeTwoStepForm - Two-step with unsubscription form (i.e. an email with confirmation link will be sent)
	UnSubscribeTwoStepForm = 3
	// UnSubscribeManual - Manual (i.e. unsubscription has to be performed by the list administrator)
	UnSubscribeManual = 4

	// FieldWizardNone .
	FieldWizardNone = "none"
	// FieldWizardFullName .
	FieldWizardFullName = "full_name"
	// FieldWizardFirstLastName .
	FieldWizardFirstLastName = "first_last_name"
)

var (
	brokenValues = []string{
		"force_subscribe",
		"require_confirmation",
		"visable",
	}

	validFieldTypes = []string{
		"text",
		"website",
		"longtext",
		"gpg",
		"number",
		"radio",
		"checkbox",
		"dropdown",
		"date-us",
		"date-eur",
		"birthday-us",
		"birthday-eur",
		"json",
		"option",
	}
)

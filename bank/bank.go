package bank

type IBank interface {
	// GetLoginMeta is used by clients to discover login criteria and dynamically
	// create login prompts
	GetLoginMeta() ILoginMeta
	// Login will submit credentials to the bank API and create a persistent session
	// used to interact and complete operations against the bank API
	Login(details ILoginDetails) (ISession, error)
}

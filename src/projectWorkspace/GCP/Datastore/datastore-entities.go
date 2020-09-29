package Datastore

/*
Assume that entities (schemas): user,Account,Services exist in dummy-namespace 
of GCP Datastore with project id "YOUR-PROJECT-ID" 

*/
func GetUserKind() string {
	return "User"
}

func GetAccountKind() string {
	return "Account"
}

func GetServicesKind() string {
	return "Services"
}

func GetNamespace() string  {
	return "dummy-namespace"	
}
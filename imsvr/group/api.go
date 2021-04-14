package group


//GroupApi 群操作相关的api
type GroupApi interface {
	 CreateGroup() error
	 RemoveGroup() error
	 RenameGroup() error
	 QuitGroup() error
	 DissolutionGroup() error
	 SetGroupAdmin() error
}



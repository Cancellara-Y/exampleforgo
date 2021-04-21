package group


//GroupApi 群操作相关的api
type Api interface {
	 //group manage
	 CreateGroup() error
	 RenameGroup() error
	 DissolutionGroup() error
	 SetGroupAdmin() error
	 GetGroupDetail() error

	 //group info
	SetGroupAvatar() error
	SetGroupName() error

	//member manage
	 RemoveMember() error
	 InviteMember() error
	 ApplyJoinGroup() error
	 QuitGroup() error

}


type Impl struct {

}
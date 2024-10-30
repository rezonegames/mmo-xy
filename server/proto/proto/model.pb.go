// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: model.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Equipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weapon   int32 `protobuf:"varint,1,opt,name=weapon,proto3" json:"weapon,omitempty"`
	Armor    int32 `protobuf:"varint,2,opt,name=armor,proto3" json:"armor,omitempty"`
	Helmet   int32 `protobuf:"varint,3,opt,name=helmet,proto3" json:"helmet,omitempty"`
	Necklace int32 `protobuf:"varint,4,opt,name=necklace,proto3" json:"necklace,omitempty"`
	Ring     int32 `protobuf:"varint,5,opt,name=ring,proto3" json:"ring,omitempty"`
	Belt     int32 `protobuf:"varint,6,opt,name=belt,proto3" json:"belt,omitempty"`
	Shoes    int32 `protobuf:"varint,7,opt,name=shoes,proto3" json:"shoes,omitempty"`
	Amulet   int32 `protobuf:"varint,8,opt,name=amulet,proto3" json:"amulet,omitempty"`
	Legguard int32 `protobuf:"varint,9,opt,name=legguard,proto3" json:"legguard,omitempty"`
}

func (x *Equipment) Reset() {
	*x = Equipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equipment) ProtoMessage() {}

func (x *Equipment) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equipment.ProtoReflect.Descriptor instead.
func (*Equipment) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *Equipment) GetWeapon() int32 {
	if x != nil {
		return x.Weapon
	}
	return 0
}

func (x *Equipment) GetArmor() int32 {
	if x != nil {
		return x.Armor
	}
	return 0
}

func (x *Equipment) GetHelmet() int32 {
	if x != nil {
		return x.Helmet
	}
	return 0
}

func (x *Equipment) GetNecklace() int32 {
	if x != nil {
		return x.Necklace
	}
	return 0
}

func (x *Equipment) GetRing() int32 {
	if x != nil {
		return x.Ring
	}
	return 0
}

func (x *Equipment) GetBelt() int32 {
	if x != nil {
		return x.Belt
	}
	return 0
}

func (x *Equipment) GetShoes() int32 {
	if x != nil {
		return x.Shoes
	}
	return 0
}

func (x *Equipment) GetAmulet() int32 {
	if x != nil {
		return x.Amulet
	}
	return 0
}

func (x *Equipment) GetLegguard() int32 {
	if x != nil {
		return x.Legguard
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Bag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Bag) Reset() {
	*x = Bag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bag) ProtoMessage() {}

func (x *Bag) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bag.ProtoReflect.Descriptor instead.
func (*Bag) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *Bag) GetItems() map[string]*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Level int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Type  string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

func (x *Skill) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Skill) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Skill) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Skills struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FightSkills map[string]*Skill `protobuf:"bytes,1,rep,name=fight_skills,json=fightSkills,proto3" json:"fight_skills,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Skills) Reset() {
	*x = Skills{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skills) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skills) ProtoMessage() {}

func (x *Skills) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skills.ProtoReflect.Descriptor instead.
func (*Skills) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{4}
}

func (x *Skills) GetFightSkills() map[string]*Skill {
	if x != nil {
		return x.FightSkills
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	KindId    int32  `protobuf:"varint,2,opt,name=kind_id,json=kindId,proto3" json:"kind_id,omitempty"`
	TaskState int32  `protobuf:"varint,3,opt,name=task_state,json=taskState,proto3" json:"task_state,omitempty"`
	StartTime int64  `protobuf:"varint,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	TaskData  string `protobuf:"bytes,5,opt,name=task_data,json=taskData,proto3" json:"task_data,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{5}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetKindId() int32 {
	if x != nil {
		return x.KindId
	}
	return 0
}

func (x *Task) GetTaskState() int32 {
	if x != nil {
		return x.TaskState
	}
	return 0
}

func (x *Task) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Task) GetTaskData() string {
	if x != nil {
		return x.TaskData
	}
	return ""
}

type Tasks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameTasks map[int32]*Task `protobuf:"bytes,1,rep,name=game_tasks,json=gameTasks,proto3" json:"game_tasks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Tasks) Reset() {
	*x = Tasks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tasks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tasks) ProtoMessage() {}

func (x *Tasks) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tasks.ProtoReflect.Descriptor instead.
func (*Tasks) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{6}
}

func (x *Tasks) GetGameTasks() map[int32]*Task {
	if x != nil {
		return x.GameTasks
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId       string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Pic          string     `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`
	RoleId       RoleType   `protobuf:"varint,4,opt,name=role_id,json=roleId,proto3,enum=proto.RoleType" json:"role_id,omitempty"`
	RoleName     string     `protobuf:"bytes,5,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	Country      string     `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	Rank         int32      `protobuf:"varint,7,opt,name=rank,proto3" json:"rank,omitempty"`
	Level        int32      `protobuf:"varint,8,opt,name=level,proto3" json:"level,omitempty"`
	Experience   int32      `protobuf:"varint,9,opt,name=experience,proto3" json:"experience,omitempty"`
	AttackValue  int32      `protobuf:"varint,10,opt,name=attack_value,json=attackValue,proto3" json:"attack_value,omitempty"`
	DefenceValue int32      `protobuf:"varint,11,opt,name=defence_value,json=defenceValue,proto3" json:"defence_value,omitempty"`
	HitRate      int32      `protobuf:"varint,12,opt,name=hit_rate,json=hitRate,proto3" json:"hit_rate,omitempty"`
	DodgeRate    int32      `protobuf:"varint,13,opt,name=dodge_rate,json=dodgeRate,proto3" json:"dodge_rate,omitempty"`
	WalkSpeed    int32      `protobuf:"varint,14,opt,name=walk_speed,json=walkSpeed,proto3" json:"walk_speed,omitempty"`
	AttackSpeed  int32      `protobuf:"varint,15,opt,name=attack_speed,json=attackSpeed,proto3" json:"attack_speed,omitempty"`
	Hp           int32      `protobuf:"varint,16,opt,name=hp,proto3" json:"hp,omitempty"`
	Mp           int32      `protobuf:"varint,17,opt,name=mp,proto3" json:"mp,omitempty"`
	MaxHp        int32      `protobuf:"varint,18,opt,name=max_hp,json=maxHp,proto3" json:"max_hp,omitempty"`
	MaxMp        int32      `protobuf:"varint,19,opt,name=max_mp,json=maxMp,proto3" json:"max_mp,omitempty"`
	AreaId       int32      `protobuf:"varint,20,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	X            int32      `protobuf:"varint,21,opt,name=x,proto3" json:"x,omitempty"`
	Y            int32      `protobuf:"varint,22,opt,name=y,proto3" json:"y,omitempty"`
	Equipment    *Equipment `protobuf:"bytes,23,opt,name=equipment,proto3" json:"equipment,omitempty"`
	Bag          *Bag       `protobuf:"bytes,24,opt,name=bag,proto3" json:"bag,omitempty"`
	Skills       *Skills    `protobuf:"bytes,25,opt,name=skills,proto3" json:"skills,omitempty"`
	Tasks        *Tasks     `protobuf:"bytes,26,opt,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *User) GetRoleId() RoleType {
	if x != nil {
		return x.RoleId
	}
	return RoleType_Mage
}

func (x *User) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *User) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *User) GetExperience() int32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *User) GetAttackValue() int32 {
	if x != nil {
		return x.AttackValue
	}
	return 0
}

func (x *User) GetDefenceValue() int32 {
	if x != nil {
		return x.DefenceValue
	}
	return 0
}

func (x *User) GetHitRate() int32 {
	if x != nil {
		return x.HitRate
	}
	return 0
}

func (x *User) GetDodgeRate() int32 {
	if x != nil {
		return x.DodgeRate
	}
	return 0
}

func (x *User) GetWalkSpeed() int32 {
	if x != nil {
		return x.WalkSpeed
	}
	return 0
}

func (x *User) GetAttackSpeed() int32 {
	if x != nil {
		return x.AttackSpeed
	}
	return 0
}

func (x *User) GetHp() int32 {
	if x != nil {
		return x.Hp
	}
	return 0
}

func (x *User) GetMp() int32 {
	if x != nil {
		return x.Mp
	}
	return 0
}

func (x *User) GetMaxHp() int32 {
	if x != nil {
		return x.MaxHp
	}
	return 0
}

func (x *User) GetMaxMp() int32 {
	if x != nil {
		return x.MaxMp
	}
	return 0
}

func (x *User) GetAreaId() int32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *User) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *User) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *User) GetEquipment() *Equipment {
	if x != nil {
		return x.Equipment
	}
	return nil
}

func (x *User) GetBag() *Bag {
	if x != nil {
		return x.Bag
	}
	return nil
}

func (x *User) GetSkills() *Skills {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *User) GetTasks() *Tasks {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x6d, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x6c, 0x6d, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x68, 0x65, 0x6c, 0x6d, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x63, 0x6b, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x63, 0x6b, 0x6c, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x65, 0x6c, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x65, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x6f, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x75, 0x6c, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x6d, 0x75, 0x6c, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x67, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x67, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x22, 0x2c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x79, 0x0a, 0x03, 0x42, 0x61, 0x67, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x61, 0x67, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x45, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x41, 0x0a,
	0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x99, 0x01, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x66,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x2e, 0x46, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x66, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x1a, 0x4c,
	0x0a, 0x10, 0x46, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8a, 0x01, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6b, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8e, 0x01, 0x0a, 0x05, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x1a,
	0x49, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd0, 0x05, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x69, 0x63, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x64, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x6f, 0x64, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x61, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x68,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6d,
	0x70, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x70, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x48, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f,
	0x6d, 0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x4d, 0x70, 0x12,
	0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x79, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x03, 0x62, 0x61, 0x67, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x67, 0x52, 0x03, 0x62,
	0x61, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x08, 0x5a,
	0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_model_proto_goTypes = []interface{}{
	(*Equipment)(nil), // 0: proto.Equipment
	(*Item)(nil),      // 1: proto.Item
	(*Bag)(nil),       // 2: proto.Bag
	(*Skill)(nil),     // 3: proto.Skill
	(*Skills)(nil),    // 4: proto.Skills
	(*Task)(nil),      // 5: proto.Task
	(*Tasks)(nil),     // 6: proto.Tasks
	(*User)(nil),      // 7: proto.User
	nil,               // 8: proto.Bag.ItemsEntry
	nil,               // 9: proto.Skills.FightSkillsEntry
	nil,               // 10: proto.Tasks.GameTasksEntry
	(RoleType)(0),     // 11: proto.RoleType
}
var file_model_proto_depIdxs = []int32{
	8,  // 0: proto.Bag.items:type_name -> proto.Bag.ItemsEntry
	9,  // 1: proto.Skills.fight_skills:type_name -> proto.Skills.FightSkillsEntry
	10, // 2: proto.Tasks.game_tasks:type_name -> proto.Tasks.GameTasksEntry
	11, // 3: proto.User.role_id:type_name -> proto.RoleType
	0,  // 4: proto.User.equipment:type_name -> proto.Equipment
	2,  // 5: proto.User.bag:type_name -> proto.Bag
	4,  // 6: proto.User.skills:type_name -> proto.Skills
	6,  // 7: proto.User.tasks:type_name -> proto.Tasks
	1,  // 8: proto.Bag.ItemsEntry.value:type_name -> proto.Item
	3,  // 9: proto.Skills.FightSkillsEntry.value:type_name -> proto.Skill
	5,  // 10: proto.Tasks.GameTasksEntry.value:type_name -> proto.Task
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	file_consts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equipment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skills); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tasks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}

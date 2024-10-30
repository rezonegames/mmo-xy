package models

import (
	"mmo-xy/pkg/z"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Equipment struct {
	Weapon   int32 `bson:"weapon"`
	Armor    int32 `bson:"armor"`
	Helmet   int32 `bson:"helmet"`
	NeckLace int32 `bson:"necklace"`
	Ring     int32 `bson:"ring"`
	Belt     int32 `bson:"belt"`
	Shoes    int32 `bson:"shoes"`
	Amulet   int32 `bson:"amulet"`
	Legguard int32 `bson:"legguard"`
}

type Item struct {
	Id    string `bson:"id"`
	Count int32  `bson:"count"`
}

type Bag struct {
	Items map[string]*Item `bson:"items"`
}

type Skill struct {
	Id    int32  `bson:"id"`
	Level int32  `bson:"level"`
	Type  string `bson:"type"`
}

type Skills struct {
	FightSkills map[string]*Skill `bson:"fight_skills"`
}

type Task struct {
	Id        int32  `bson:"id"`
	KindId    int32  `bson:"kind_id"`
	TaskState int32  `bson:"task_state"`
	StartTime int64  `bson:"start_time"`
	TaskData  string `bson:"task_data"`
}

type Tasks struct {
	GameTasks map[int32]*Task `bson:"game_tasks"`
}

type User struct {
	Name         string `bson:"name"`
	UserId       string `bson:"_id"`
	Pic          string `bson:"pic"`
	RoleId       int32  `bson:"role_id"`
	RoleName     string `bson:"role_name"`
	Country      string `bson:"country"`
	Rank         int32  `bson:"rank"`
	Level        int32  `bson:"level"`
	Experience   int32  `bson:"experience"`
	AttackValue  int32  `bson:"attack_value"`
	DefenceValue int32  `bson:"defence_value"`
	HitRate      int32  `bson:"hit_rate"`
	DodgeRate    int32  `bson:"dodge_rate"`
	WalkSpeed    int32  `bson:"walk_speed"`
	AttackSpeed  int32  `bson:"attack_speed"`
	Hp           int32  `bson:"hp"`
	Mp           int32  `bson:"mp"`
	MaxHp        int32  `bson:"max_hp"`
	MaxMp        int32  `bson:"max_mp"`
	AreaId       int32  `bson:"area_id"`
	X            int32  `bson:"x"`
	Y            int32  `bson:"y"`
	*Equipment   `bson:"equipment"`
	*Bag         `bson:"bag"`
	*Skills      `bson:"skills"`
	*Tasks       `bson:"tasks"`
}

func (p *User) ConvToProto() *proto.User {
	var (
		itemList = make([]*proto.ItemInfo, 0)
	)

	return &proto.User{}
}

func GetUser(userId string, fields ...string) (*User, error) {
	var (
		filter = bson.M{
			"_id": userId,
		}
		opts       = &options.FindOneOptions{}
		projection = bson.M{}
		p          = &User{}
		err        error
	)

	for _, field := range fields {
		projection[field] = 1
	}
	opts.Projection = projection

	err = mClient.FindOne(p, DB_NAME, COL_PLAYER, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, z.NilError{Msg: err.Error()}
		}
		return nil, err
	}
	return p, nil
}

func AddItems(userId string, itemList []*proto.ItemInfo) error {
	var (
		filter = bson.M{
			"_id": userId,
		}
		update = bson.D{}
	)
	return mClient.UpsertOne(DB_NAME, COL_PLAYER, filter, update)
}

func CreateUser(name string, pic string, coin int32, roleId int32, country string) (*User, error) {
	var (
		err      error
		user     *User
		userId   string
		insertId any
	)

	userId, err = GetNextID(COL_PLAYER)
	if err != nil {
		return nil, err
	}
	// cmd+.
	user = &User{Name: name, UserId: userId, Pic: pic, RoleId: roleId, RoleName: "", Country: country, Rank: 0, Level: 0, Experience: 0, AttackValue: 0, DefenceValue: 0, HitRate: 0, DodgeRate: 0, WalkSpeed: 0, AttackSpeed: 0, Hp: 0, Mp: 0, MaxHp: 0, MaxMp: 0, AreaId: 0, X: 0, Y: 0,
		Equipment: &Equipment{
			Weapon:   0,
			Armor:    0,
			Helmet:   0,
			NeckLace: 0,
			Ring:     0,
			Belt:     0,
			Shoes:    0,
			Amulet:   0,
			Legguard: 0,
		}, Bag: &Bag{
			Items: map[string]*Item{},
		}, Skills: &Skills{
			FightSkills: map[string]*Skill{},
		}, Tasks: &Tasks{
			GameTasks: map[int32]*Task{},
		}}

	insertId, err = mClient.InsertOne(DB_NAME, COL_PLAYER, user)
	user.UserId = insertId.(string)
	return user, err
}

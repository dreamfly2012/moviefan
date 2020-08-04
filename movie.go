package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Config struct {
	Mongo struct {
		Addrs      string `mapstructure:"host"`
		Database   string `mapstructure:"db"`
		Port       int64  `mapstructure:"port"`
		Username   string `mapstructure:"username"`
		Password   string `mapstructure:"password"`
		Maxconnect int64  `mapstructure:"maxconnect"`
	}
}

func connect(cName string) (*mgo.Session, *mgo.Collection) {
	//c := &MongoDialInfo{}
	var c Config
	file := "./config.yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config error:", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("unmarshal config failed", err)
	}
	fmt.Printf("movie的结构%v", c)
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/?authSource=admin", c.Mongo.Username, c.Mongo.Password, c.Mongo.Addrs, c.Mongo.Port)
	fmt.Println(url)
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session, session.DB(c.Mongo.Database).C(cName)
}

//Movie 定义电影数据结构
type Movie struct {
	ID           bson.ObjectId `bson:"_id"`
	IDS          int32         `bson:"id"`
	Mid          int32         `bson:"mid"`
	Name         string        `bson:"name"`
	Summary      string        `bson:"summary"`
	Director     string        `bson:"director"`
	Actor        string        `bson:"actor"`
	Type         string        `bson:"type"`
	Region       string        `bson:"region"`
	Publishdate  int32         `bson:"publishdate"`
	Avgrating    float32       `bson:"avgrating"`
	Commentcount int32         `bson:"commentcount"`
	Ratingcount  int32         `bson:"ratingcount"`
	Addtime      string        `bson:"addtime"`
	Updatetime   string        `bson:"updatetime"`
}

func (a *Movie) save() error {
	s, c := connect("movie")
	defer s.Close()
	return c.Insert(&a)
}

func (a Movie) all() ([]Movie, error) {
	s, c := connect("db_movie")
	defer s.Close()
	var group []Movie
	err := c.Find(nil).All(&group)
	return group, err
}

func (a *Movie) random() Movie {
	s, c := connect("db_movie")
	defer s.Close()

	pipeline := []bson.M{
		bson.M{"$sample": bson.M{"size": 1}},
	}

	pipe := c.Pipe(pipeline)

	iter := pipe.Iter()

	for iter.Next(&a) {

	}
	return *a

}

func (a *Movie) get(name string) error {
	s, c := connect("db_movie")
	defer s.Close()
	err := c.Find(bson.M{"name": name}).One(&a)
	if err != nil {
		fmt.Println("mongo get error", err)
	}
	fmt.Println(a)
	return err
}

func (a *Movie) delete(id bson.ObjectId) error {
	s, c := connect("movie")
	defer s.Close()
	return c.Remove(bson.M{"_id": id})
}

func (a *Movie) update() error {
	s, c := connect("movie")
	defer s.Close()
	c.Update(bson.M{"_id": a.Name}, a)
	return a.get(a.Name)
}

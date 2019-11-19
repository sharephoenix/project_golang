package main

func main() {
}


//func case2() {
//
//	// Creating UUID Version 4
//	// panic on error
//	u1 := uuid.Must(uuid.NewV4())
//	fmt.Printf("UUIDv4: %s\n", u1)
//
//	// or error handling
//	u2, err := uuid.NewV4()
//	if err != nil {
//		fmt.Printf("Something went wrong: %s", err)
//		return
//	}
//
//	token := GetMd5String(base64.URLEncoding.EncodeToString(u2.Bytes()))
//
//	fmt.Printf("UUIDv4: %v\n", token)
//
//	// Parsing UUID from string input
//	//u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
//	//if err != nil {
//	//	fmt.Printf("Something went wrong: %s", err)
//	//	return
//	//}
//	//fmt.Printf("Successfully parsed: %s", u2)
//}
//
////生成32位md5字串
//func GetMd5String(s string) string {
//	h := md5.New()
//	h.Write([]byte(s))
//	return hex.EncodeToString(h.Sum(nil))
//}
//
////生成Guid字串
//func UniqueId() string {
//	b := make([]byte, 48)
//
//	if _, err := io.ReadFull(rand.Reader, b); err != nil {
//		return ""
//	}
//	return GetMd5String(base64.URLEncoding.EncodeToString(b))
//}

/************************************/

//func case1 {
	//
	//c, err := redis.Dial("tcp", "127.0.0.1:8001")
	//if err != nil {
	//	fmt.Println("Connect to redis error", err)
	//	return
	//}
	//defer c.Close()
	//
	//_, err = c.Do("SET", "mykey", "superWang")
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}
	//
	//username, err := redis.String(c.Do("GET", "mykey"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//} else {
	//	fmt.Printf("Get mykey: %v \n", username)
	//}
//}

/************************************/

//func case0() {
//	fmt.Println("this is redis")
//
//	client := redis.NewClient(&redis.Options{
//		Addr:     "localhost:8001",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//
//	defer fmt.Println("THIS REDIS END")
//	defer client.Close()
//
//	pong, err := client.Ping().Result()
//	if err == nil {
//		fmt.Println("链接成功: %v", pong)
//		Set(client)
//		Get(client)
//		return
//	}
//
//	fmt.Println("链接失败：%v", err.Error())
//}
//
//func Set(client *redis.Client) {
//	user := User{
//		"alexluan",
//		"anhui",
//		1,
//	}
//	jsonBytes, err := json.Marshal(user)
//	if err != nil {
//		fmt.Println(err)
//	}
//	err0 := client.Set("user", string(jsonBytes), 0).Err()
//	if err0 != nil {
//		panic(err)
//	}
//
//	val, err := client.Get("user").Result()
//	if err == nil {
//		var userr User
//		e := json.Unmarshal([]byte(val), &userr)
//		if e == nil {
//			fmt.Println("struct", userr)
//		} else {
//			fmt.Println("errrrror", e.Error())
//		}
//	}
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("user", val)
//}
//
//func Get(client *redis.Client) {
//	val2, err := client.Get("key").Result()
//	if err == redis.Nil {
//		fmt.Println("key does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key", val2)
//	}
//}
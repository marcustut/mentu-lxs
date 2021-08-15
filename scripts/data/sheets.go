package main

import (
	"context"
	"fmt"
	// "cloud.google.com/go/firestore"
	// "github.com/kr/pretty"
)

func main() {
	users := map[interface{}]interface{}{
		"Wong Hao Jun":           "1KANcTSdrtRFFTVylojrgrXcfRO2",
		"Kristin Pek Su-Wen":     "xM1Flb868aVoIRcrR2mTO8MLLnG3",
		"Fong Jia Shin":          "4j5P8vBfrFXgXvZ1jVJMkXDUqSy2",
		"Chow Kian Lam":          "Q158kqdTIIMCbeU63vI91nJmAMX2",
		"Chew Tse Wen":           "jV2z8sNSjFVti8NHjm5TX2IdEHf2",
		"Goh Hui Wen":            "LLXvCJD64UYvt2a2uX9NKqK9De23",
		"Cynthia Ding Chiau Wei": "kR0Bax9L01bwWLGjxuLQazuqjGi2",
		"Andre Him Jian Zhi":     "Nc1Q2nSjkdTlsph5WoHy6QzWCft1",
		"Justin Shuck Chze Wei":  "nNevhkXDQDOCfmRUIMrv5rqbg9G3",
		"Yee Sook Fun":           "V9nKuGWcUMNZNeAWa6BDyvo3pOo2",
		"Ong Sze Miq":            "8SDBD0CpivPZvUHPo6do2ep5rby1",
		"Li Kar Ken":             "zOjv5W4FSSTTZXujhVb7FdumXKm1",
		"Lee Ting Chung":         "CwrwyCY9DAamykjszIR25ZG10SC3",
		"Yee Chee How":           "asVe7b0UpUXKC4E9gEe670GMTSy2",
		"Khoo Han Sheng":         "vhKJeIOk9LfRspvwTb3YzIpsdkY2",
		"Liew Yu Jing":           "0QkmJHE2AfTIYwFf4467rmMqIGE2",
		"Chia Chee Hong":         "qBXuVstfKOVQjhy4VHl5rqJSY3W2",
		"Lim Shu Yi":             "Z5MFBUWjN8aZnnb6m9DTroou6rb2",
		"Chai Jun Wai, William":  "pP4YU1eSguVzZ9cK7yzUWnCf1di1",
		"Kelvin Lim An Hong":     "uhOvbAWX03XDDoyZBxBwLAvX9Zi2",
		"Myona Hong Sook Cheng":  "4MiTP979R1URzA0ZJ3x8UCsQTvw2",
		"Ng Mei Mei":             "ewJYUqY6YufButsgoxeB0pLhKax1",
		"Foo Sze Yun":            "TCNFImSjpXd9HkPKNyBa3pchymf1",
		"Lim Xin Yee":            "O0DK8X72FUYoIabCPmP2GV19R553",
		"Wong Yui Soon":          "iVYGxL7qBxNIdPbr2APkiEaUviL2",
		"Kenneth Lee Kar Hong":   "n90UkGLGfCRp2JNM3Nnvh2gXR253",
		"Christy Ling Xin Yan":   "EA4azazjP5eloWPmcTYdGennLsN2",
		"Chow Hwee Lu":           "hebs0KFGHYhBnwROQHCPkl0cYhB2",
		"Yong Shee Mun (Kate)":   "rIkMGZ9suWSZRNRZbofbpYEpbMz2",
		"Chin Pei Yueh":          "P3GfUU6b9wNaFsIRlw6N2Zgv8kJ2",
		"Ng Kok Cheng":           "WOFWxS0EjVdxedskYRQrxCy76gX2",
		"Lim Mei Yee":            "3D4fKDPlyfcKo4CTCCHGm9c4hVk2",
		"Loo Kai Lee":            "xQB9q8AoMIYcaEOmDLk6hdxa4tH3",
		"Yap Jia Qian":           "Y0qGe7CZBCd8w3oTxhmkPtzFsFE3",
		"Kok Yi Cheng":           "eH3h7SRm29Ne04t9yQK1Ui4fDHv1",
		"Toh Kar Wei":            "RUCUKmyjIjS6XXilu9yEsMueyD52",
		"Nicole Teo Shu Ern":     "VTOCF0RWBTXoQzKljCkRXAsmRqE3",
		"Too Yu Han":             "Ybk67WUXQFU41upbc6zskqGjm852",
		"Yap Pik Xia":            "o8TdVayAq7W1QxrrFnVWpNsmFY92",
		"Nai Pui Shan":           "eIJjhS0fgITiEjqGHqEJD5s5PTI3",
		"Pang Zheng Yeong":       "mjjlWYE8g7MYJM7sM9a6Gg5ZFRJ3",
		"Lee Mei Kee":            "nYNMuJtvOJNsMSnvsjO7GdYdcWb2",
		"Soh Siau Po":            "xYFumrVp30fyYi96PYayWqgaH6I2",
		"Lim Chen Tung":          "lBKvk38AR2fuPs5JNNaVSGO9UFJ3",
		"Lim Ming Zhou":          "oEaDRGC2c7e7IzI7NxbwJ6KXcij2",
		"Lim Xin Yi":             "04KQzWgLNDXiXeUIV0zPojQWTPn2",
		"Goh Wei Heng":           "B8ZsneX2goMS3GZduMeIwcmqkfe2",
		"Lim Zu Yi":              "WKxG4X2LRwVejgDCZy65fAaNhys2",
		"Chang Xiang Ying":       "n43LaQMZHXZgEuShlduGbzrJ8X53",
		"Chan Yung You":          "sq5oKZlctCN6Uytt1mpzkfu6rzq1",
		"Lim Xin Ru":             "9i1QD6H9BiO1UyDrlT4bxPyREgg2",
		"Wong Jien Thong":        "YkSoWIUVDpaIwkYpm3k5VNsDZGB3",
		"Chan Mun Lik":           "Xv3puObpa7VbCeZFS7KrB5MkbaH2",
		"Wong Kar Shean":         "tLaXYrMZDtP1Vp7xLUV80oov7YF3",
		"Jenny Then Tzu Ying":    "npSw4EC6aZSVVkt4YQvQzE2qxXz2",
		"Goh Cheng Hao":          "mv8nlGz5qJaeToMYZ8EXdWviRP83",
		"Law Jia Yang":           "Xv3puObpa7VbCeZFS7KrB5MkbaH2",
	}

	ctx := context.Background()

	ss, err := NewSpreadSheetService("/client_secret.json")
	if err != nil {
		panic(fmt.Sprintf("Unable to start spreadsheet service: %v", err))
	}

	values, err := ss.Read("1rGCfteAxl7fxcwcwEC-q2DlWcFDrXD5mEbqXVEdElDw", "Data!A2:J")
	if err != nil {
		panic(fmt.Sprintf("Error reading data from spreadsheet: %v", err))
	}

	fs, err := NewFirebaseService("/client_secret.json")
	if err != nil {
		panic(fmt.Sprintf("Error initializing firebase: %v", err))
	}

	db, err := fs.app.Firestore(ctx)
	if err != nil {
		panic(fmt.Sprintf("Error initializing firestore: %v", err))
	}
	defer db.Close()

	ref := db.Collection("reports")

	for _, value := range values {
		ref.NewDoc().Set(ctx, map[string]interface{}{
			"uid":  users[value[0]],
			"name": value[0],
			"scores": map[string]interface{}{
				"time":           value[1],
				"communication":  value[2],
				"willingness":    value[3],
				"responsibility": value[4],
				"maturity":       value[5],
			},
			"strength":    value[6],
			"words":       value[7],
			"total":       value[8],
			"grand_total": value[9],
		})
	}
}

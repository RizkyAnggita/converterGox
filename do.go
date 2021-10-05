package converterGox

import (
	"regexp"
	"strings"
)

func ConvertJSONGraphQL(jsonString string) (string, error){
    regex, err := regexp.Compile(`\"([a-zA-Z0-9_]+)\":`)
    if err != nil{
      return "", err
    }

    temp := jsonString
    // str := "\"customerRequestId\": \"hehe\", \"customerid\": \"hege\"" 

    if (regex.MatchString(temp)){
      replacer := regex.FindAllStringSubmatch(temp, -1)
      for i := 0; i < len(replacer); i++ {
        temp = strings.Replace(temp, replacer[i][0], replacer[i][1]+": " , -1 )
      }
    }
    return temp, nil
}

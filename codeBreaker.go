package main

var secretCode string = "1234"
func codebreaker(code string) string {

  var answer string
  // var elemento1 string = code[0:len(code)-3]
  // var elemento2 string = code[1:len(code)-2]
  // var elemento3 string = code[2:len(code)-1]
  // var elemento4 string = code[3:len(code)]
  //
  // var secret1 string = secretCode[0:len(secretCode)-3]
  // var secret2 string = secretCode[1:len(secretCode)-2]
  // var secret3 string = secretCode[2:len(secretCode)-1]
  // var secret4 string = secretCode[3:len(secretCode)]

  // if (elemento1 == secret1 && elemento2 == secret2 && elemento3 == secret3 && elemento4 == secret4) {
  //   answer = "XXXX"
  // }
  // if (elemento1 == secret1 || elemento1 == secret2 || elemento1 == secret3 || elemento4 == secret4) {
  //   answer = "XXXX"
  // }
  count_underline := 0;
  for i := 0; i < len(secretCode); i++ {
    if (secretCode[i]==code[i]){
      answer = answer + "X"
    }else{

    for j :=0; j < len(code); j++{


      if(code[i]==secretCode[j]){
        count_underline = count_underline + 1

      }
      if (count_underline > 0){
        answer = answer + "_"
        count_underline = 0
      }
    }
    }
  }

return answer
}

func setSecret(secret string){
  secretCode = secret
}

 package pl.edu.uj.demo
 import org.springframework.http.HttpStatus
 import org.springframework.web.bind.annotation.GetMapping
 import org.springframework.web.bind.annotation.RequestMapping
 import org.springframework.web.bind.annotation.RestController
import org.springframework.stereotype.Component;
import org.springframework.beans.factory.annotation.Autowired;

 import javax.servlet.http.HttpServletRequest

 @Service
 class LoginService(){
  @Autowired
     fun zaloguj(login: String, password: String) {
     val haslo="haslo"
     val log="log"
     if (login==log && password==haslo){
       println("Zalogowano!")
     }
     else{
       println("Podano nieprawidlowe dane logowania")
     }
 }
}

 @RestController
 @RequestMapping("/product")
 class ProductController{
          @GetMapping("/test")
          fun testowy()="OK"
 }


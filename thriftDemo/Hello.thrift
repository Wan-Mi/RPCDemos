

struct User {
    1: required string name
    2: required i32 age
}

service HelloService {
    User sayHello(1: string userName, 2:i32 userAge)
}
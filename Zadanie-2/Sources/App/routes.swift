import Fluent
import Vapor

func routes(_ app: Application) throws {
    app.get { req in
        return "It works!"
    }

    app.get("hello") { req -> String in
        return "Hello, world!"
    }

    try app.register(collection: MuzykController())
    try app.register(collection: GrupaaController())
    try app.register(collection: OrkiestraController())
}
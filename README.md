# graphql-twitter
follow this tutorial: 
https://store.equimper.com/view/courses/build-a-twitter-clone-graphql-api-using-golang/765944-introduction/2238834-go-modules

# Folder structure
- cmd  
root folder which is our program start entrypoint
- config  
app configuration globally accessible for our app like `database url`, `JWT secret` etc.
- domain  
  business logic inside this folder  
  - domain use interface to do it logic, decoupling the implementation from repo folder, easy to mock unit test   
  - domain also implemented the graphql interface, decoupling the implementation from graphql-handler  
  - graphql logic for handler to call, we don't want to add all logic inside our graphql layer.  
  We create the domain layer to encapsulate all functions for graphql-handler to call.  
  This will let you easily switch from graphql to restful api in the future.  
- postgres  
All repo function talk to database will place in here.  
Create MongoDB, MySql folder if you need to communicate with different dbs
- auth.go, user.go
All Interface place in here which will be used by repo layer or domain layer    
We also have input Validation and Sanitize function in here for repo or domain layer to use.  
This way to prevent circular dependency.  
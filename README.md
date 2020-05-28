# Team-Members
Team-Members is a back-end application that helps organizations manage their teams. It consists of a Go RESTful API app and uses [MongoDb](https://www.mongodb.com/) for the database. The [Gin Web Framework](https://github.com/gin-gonic/gin) package is used for HTTP routing and the [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) package is used to connect and work with MongoDB.

# Features and Requirements

- A member has a name and a type. The member type can be an employee or a contractor - if it's a contractor, the duration of the contract needs to be saved and if it's an employee we need to store their role, for instance: Software Engineer, Project Manager and so on.

- A member can be tagged, for instance: C#, Angular, General Frontend, Seasoned Leader and so on. (Tags will likely be used as filters later, so keep that in mind)

A RESTful CRUD is offered for the information above.

# Installation
#### 1. Install Go
Follow the instructions here: https://golang.org/doc/install

#### 2. Install MongoDB
Follow the instructions here: https://docs.mongodb.com/manual/installation/

#### 3. Install Gin package
```bash
go get -u github.com/gin-gonic/gin
```
#### 4. Install the MongoDB Go Driver package
```bash
go get github.com/mongodb/mongo-go-driver
```
NOTE: The output of this may look like a warning stating something like package `github.com/mongodb/mongo-go-driver`: no Go files in (...). This is expected output.

#### 5. Install [gotoenv](https://github.com/subosito/gotenv) package
```bash
go get github.com/subosito/gotenv
```

#### 6. Install Team-Members
Copy team-members directory to /go/src


NOTE: Examples shown here below use the default configuration settings avaiable in the .env file under team-members directory. Feel free to change the settings to your preferred configuration settings:

```bash
    MONGODB=mongodb://localhost:27017
    DB_NAME=members

    #PORT=8080
```
PORT environment variable is for the Gin engine and if not set team-members/main.go will default it to "8080"


## To Run
1. Start MongoDB
    ```bash
    sudo mongod
    ```
3. Build team-members. From /go/src/team-members run
    ```bash
    go build
    ```
4. Run team-membrers. From /go/src/team-members run 
    ```bash
    go run team-members
    ```

## Usage and Workflow
#### NOTE: I highly recommend using [Postman](https://www.getpostman.com/) to work with and test Team-Members RESTful API Golang/MongoDB app.
1. #### Create member
    The payload data to create new employee and contractor members through a POST request will be the following JSON data structures respectively:

    Employee:
    ```json
    {
        "oid": null,
        "name": "Joel McDaniel",
        "type": {
            "key": 1,
            "value": "Employee",
            "properties": {
                "role": "Senior Golang Engineer"
            }
        },
        "tags": [
            "Go",
            "Golang",
            "MongoDB"
        ]
    }
    ```

    Contractor:
    ```json
    {
        "oid": null,
        "name": "Someone Else",
        "type": {
            "key": 0,
            "value": "Contractor",
            "properties": {
                "duration": "6 months"
            }
        },
        "tags": [
            "JavaScript",
            "HTML5/CSS3",
            "Vue.js"
        ]
    }
    ``` 
    
    NOTE: For your convenience the above JSON data can be also be found under the team-members directory in the create_member_payload_data.json file.

    Using Postman (or some similar tool) use the member json data structures above for the body of the POST request to be posted as raw JSON.
    
    To create a new member in the database POST the new member JSON to /api/member like so: 

        POST: http://localhost:8080/api/member

    For the first member, this will create the mongodb "members" database and insert the document record.

2. #### Read all members
    To read the collection of members inserted into the "members" database use Postman to make a GET request to the same URI as above like so:

        GET: http://localhost:8080/api/members

    
    The existing members collection response will be returned as JSON and each members' oid property value will be set to the document record ObjectID string value. This is the unique primary key value MongoDB generated during the inserts represented as the _id field in the "members" database. This value will be used in the request URI to fetch individual members as shown later below in section 4.

    ```json
    [
        {
            "oid": "5ce9f998527c85b1d5a930f1",
            "name": "Joel McDaniel",
            "type": {
                "key": 1,
                "value": "Employee",
                "properties": {
                    "role": "Senior Golang Engineer"
                }
            },
            "tags": [
                "Go",
                "Golang",
                "MongoDB"
            ]
        },
        {
            "oid": "5ce9f9c4527c85b1d5a930f2",
            "name": "Someone Else",
            "type": {
                "key": 0,
                "value": "Contractor",
                "properties": {
                    "duration": "6 months"
                }
            },
            "tags": [
                "JavaScript",
                "HTML5/CSS3",
                "Vue.js"
            ]
        }
    ] 
    ```
3. #### Read members by tags
    To read members filtered by tags make a get request to /api/members/?tags=`<tag>`&tags=`<tag>` ect. appending /? and tags querystring parameter for each tag you wish to filter on like so:

            GET: http://localhost:8080/api/members/?tags=Meowin&tags=Hunting&tags=Go

    This will return members corresponding to the tags filtered on from the request URI querystring:

    ```json
    {
        "oid": "5ce9f998527c85b1d5a930f1",
        "name": "Joel McDaniel",
        "type": {
            "key": 1,
            "value": "Employee",
            "properties": {
                "role": "Senior Golang Engineer"
            }
        },
        "tags": [
            "Go",
            "Golang",
            "MongoDB"
        ]
    },
    {
        "oid": "5cec7a92fc1f14ada5542358",
        "name": "Franz Wooleybear",
        "type": {
            "key": 1,
            "value": "Employee",
            "properties": {
                "role": "Señor Kitty Cat"
            }
        },
        "tags": [
            "Meowin",
            "Loungin",
            "Hunting"
        ]
    }
    ```

4. #### Read member

    By ObjectID:

    To read individual existing members by ObjectID use Postman to make a GET request to /api/member/oid/:oid appending the oid string value to the URI endpoint path like so:

            GET: http://localhost:8080/api/member/oid/5ce9f998527c85b1d5a930f1

    This will return the member corresponding to the string ObjectID (:oid):

    ```json
    {
        "oid": "5ce9f998527c85b1d5a930f1",
        "name": "Joel McDaniel",
        "type": {
            "key": 1,
            "value": "Employee",
            "properties": {
                "role": "Senior Golang Engineer"
            }
        },
        "tags": [
            "Go",
            "Golang",
            "MongoDB"
        ]
    }
    ```
    By Name:

    To read individual existing members by Name use Postman to make a GET request to /api/member/name/:name appending the name string value to the URI endpoint path like so:

            GET: http://localhost:8080/api/member/name/Franz Wooleybear

    This will return the member corresponding to the string name (:name):

    ```json
    {
        "oid": "5cec7a92fc1f14ada5542358",
        "name": "Franz Wooleybear",
        "type": {
            "key": 1,
            "value": "Employee",
            "properties": {
                "role": "Señor Kitty Cat"
            }
        },
        "tags": [
            "Meowin",
            "Loungin",
            "Hunting"
        ]
    }
    ```


5. #### Update member
    To update a member use Postman with the JSON member data from a GET request (including oid value) for the body of the PUT request as raw JSON making any changes to the data before sending the request.
    
    ```json
    {
        "oid": "5ce9f9c4527c85b1d5a930f2",
        "name": "Someone Else",
        "type": {
            "key": 1,
            "value": "Employee",
            "properties": {
                "role": "Front-end Developer"
            }
        },
        "tags": [
            "JavaScript",
            "HTML5/CSS3",
            "Vue.js"
        ]
    }
    ```

     Make the PUT request to /api/member/:oid with the oid string value appended to the URI endpoint path like so:

        PUT: http://localhost:8080/api/member/5ce9f9c4527c85b1d5a930f2

    Make another GET request at the same URI to verify/return the updated document record.

6. #### Delete member
    To delete a member use Postman to make a DELETE request to /api/member/:oid with the oid string value appended to the URI endpoint path like so:

        DELETE: http://localhost:8080/api/member/5ce9f9c4527c85b1d5a930f2

    Make another GET all members request to verify the member was removed from the collection.

## Hope you enjoyed!

    

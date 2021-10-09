# Instagram-Backend-API
This is the task done for Appointy 

Tasks Done here : 

		Create an User
				Should be a POST request
				Use JSON request body	
				URL should be ‘/users'
				
		Get a user using id
				Should be a GET request
				Id should be in the url parameter
				URL should be ‘/users/<id here>’
				
I used Go for the very first time, i have used other backend languages earlier but this one was quite different. Documentation provided on the internet is very easy to grasp and move forward with your project. I

Basic Setup :

		go get github.com/mongodb/mongo-go-driver

Connection of Go with MongoDb :

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

I will work upon this project further and will update the Repo in future. 

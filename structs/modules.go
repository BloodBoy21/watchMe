package structs

type CommandsRun struct {
	Name 			string `short:"n" long:"name" description:"Name of the service"`
	DockerFile 		string `short:"d" long:"dockerfile" description:"Dockerfile to use"`
	Dockerize 		bool   `short:"z" long:"dockerize" description:"Dockerize the service"`
	CodeLang 		string `short:"l" long:"lang" description:"Programming language to use"`
	Args  []string `short:"a" long:"args" descriptions:"Arguments to pass to docker run"`
}

type CommandsInit struct {
	
}

Command1 -1:
protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto     
Command1 -2:
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/yuvakkrishnan/grpc-go --go-grpc_out=. --go-grpc_opt=module=github.com/yuvakkrishnan/grpc-go greet/proto/dummy.proto


Create a Makefile in root level....
Installation of Makefile...

Run commands => make greet
Run command => make help =>
all                                           Generate Pbs and build                                                                    greet                                         Generate Pbs and build for greet                                                          calculator                                    Generate Pbs and build for calculator                                                     blog                                          Generate Pbs and build for blog                                                           
test                                          Launch tests
clean                                         Clean generated files
clean_greet                                   Clean generated files for greet      
clean_calculator                              Clean generated files for calculator 
clean_blog                                    Clean generated files for blog       
rebuild                                       Rebuild the whole project
bump                                          Update packages version
about                                         Display info related to the build
help                                          Show this help



Makefile (Windows)
In order to make the build more replicable, I'm using make in this section. In order to install that, I recommend you install Chocolatey (https://chocolatey.org/install), and then run:

choco install make
You should then be able to use make command in the directory that contains the Makefile.



This is optional, since you can still build the .proto files by hand by running the following commands:

protoc -I${PROJECT}/proto --go_opt=module=${YOUR_MODULE} --go_out=. ${PROJECT}/proto/*.proto
where ${YOUR_MODULE} is the name of your go module (you can find that in your go.mod file) and ${PROJECT} is one of the projects name (greet, calculator, blog).
## Purpose

Exercise my go skills, build something fun that I can use every day. I have to
join a zoom session every day. Why do I have to click a link? That is annoying
to do every day. So why not automate it?




## Features
* Makefile to build consistently in a local environment and remote environment
* Dockerfile for a generic image to build for 
* Go Mod (which you should to your project path change)
* VS Code environment
* Generic docker push

## TODO 
* build a grpc service that opens up a zoom session
* build a grpc service that reads the calendar
* build a sqllite storage service that stores state
* build an event source system to detect the change and send a grpc message to
  the correct server? -- overkill but fun
* experiment with running services in different locations. Containers locally
  and Containers remotely or in Lambda things because why not, interesting to
learn.
* 


## Installing via brew
* `brew install --verbose --build-from-source brew/Formula/go-grpc-zoom-manager.rb`

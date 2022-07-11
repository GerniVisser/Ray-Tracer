# Ray Tracing  - Golang 
First implementation of a simple Ray-Tracer in Golang. This implementation is heavily inspired by the excellent book by Jamis Buck, [The Ray Tracer Challenge](https://pragprog.com/book/jbtracer/the-ray-tracer-challenge). It is a complete explanation of how a Ray Tracer work and can be implemented in any language of your choice. 

![Scene-_2](https://user-images.githubusercontent.com/46093495/178142088-569fe5e8-3312-4041-9302-e14aa6202799.png)

## About the project
I wanted to learn a low level language for a while and decided this project would be the perfect opportunity. I first thought of using C++ when I discovered [Go](https://go.dev/) a low level language developed by Google with just 27 keywords. In retrospective I think using C++ or Rust would have been better options for this particular project. 

**Goals**

* Build a simple fully functional Ray-Tracer 
* Not use any external packages

## Usage 
-   Install  [Go](https://go.dev/) 
- Clone repo to local repository 
	- `cd <directory>`
	- `git clone https://github.com/GerniVisser/Ray-Tracer.git`
- In `main.go` set the scene you want to render.
> 2 Default scenes pre-loaded 
- To render a custom scene create new go file in **scene** folder
	- follow conventions in **scene1.go** and **scene2.go**
	- In **main.go** set you new scene
- To render new image run `go run main.go`

## Features 
### Render and reshape Spheres and Planes
![Reshape](https://user-images.githubusercontent.com/46093495/178223761-401c604b-d5b0-4fb0-8033-a75c62422cf6.png)

### Reflections and Shadows
![Reflection-Demo](https://user-images.githubusercontent.com/46093495/178142176-c347f2d8-4ed3-4cd1-a315-b54b3df38408.png)

### Refraction
![Refraction](https://user-images.githubusercontent.com/46093495/178142312-90333f4b-b977-4724-b0db-4914fe44c561.png)

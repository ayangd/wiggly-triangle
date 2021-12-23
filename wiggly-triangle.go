package main

import (
	_ "image/png"
	"log"
	"math"

	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	phase     float32
	rotationY float32
)

const width, height = 800, 600

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(width, height, "Wiggly Triangle", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	setupScene()
	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
		glfw.SwapInterval(1)
	}
}

func setupScene() {
	gl.Enable(gl.DEPTH_TEST)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	f := ((float64(width) / height) - 1) / 2
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -3.0)
	phase += 0.05
	if phase > 3.1415*2 {
		phase -= 3.1415 * 2
	}
	rotationY += 0.5
	if rotationY > 360 {
		rotationY -= 360
	}
	gl.Rotatef(rotationY, 0, 1, 0)

	gl.Begin(gl.TRIANGLES)

	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, -1, Sin32(phase))
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(1, -1, Sin32(phase+(3.1415*2.0/3.0)))
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 1, Sin32(phase+(3.1415*2.0/(2.0/3.0))))

	gl.End()

	gl.Begin(gl.LINES)

	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, -1, -1)

	gl.Color3f(0, 1, 0)
	gl.Vertex3f(1, -1, 1)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(1, -1, -1)

	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 1, 1)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 1, -1)

	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(1, -1, 1)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(1, -1, -1)

	gl.Color3f(0, 1, 0)
	gl.Vertex3f(1, -1, 1)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 1, 1)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(1, -1, -1)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 1, -1)

	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 1, 1)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 1, -1)

	gl.End()
}

func Sin32(rad float32) float32 {
	return float32(math.Sin(float64(rad)))
}

func Cos32(rad float32) float32 {
	return float32(math.Cos(float64(rad)))
}

package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jcow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 256H768q0 192-14 316t-46.5 196T629 868t-117 28q-64 0-123-17t-96-47t-37-64q0-26 .5-39t3.5-35t9.5-35t19.5-27.5t31-23.5q-6-27-25.5-87.5t-29-119.5t-9.5-145H128q-61 0-94.5-13.5T0 192q0-43 39-85.5T128 64q29 0 67.5 15.5T266 119q22-52 90-85.5T512 0t156 33.5t90 85.5q32-24 70.5-39.5T896 64q50 0 89 42.5t39 85.5q0 37-33.5 50.5T896 256zM480 64q-32 0-52 4.5T399 83t-12 19.5t-3 25.5q0 19 10 43t22 52t22 104t10 185q0 42-10.5 81.5T416 640q-96 47-96 96q0 43 64.5 69.5T533 832q5-4 12-11.5t19-32.5t12-52q0-99 16-242t32-250t16-116q0-27-47-45.5T480 64z"/>`),
		g.Group(children),
	)
}
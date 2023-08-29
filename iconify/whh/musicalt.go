package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 416q0 60-34.5 129.5t-80 114T576 704q-25 0-44.5-20T512 640q0-26 18.5-45t45.5-19q-1-8-4-21t-14-46t-24-58t-36-46t-50-21q-35 0-49.5 13.5T384 448v448q0 53-56 90.5T192 1024T56 986.5T0 896t56-90.5T192 768q31 0 64 8V64q0-26 19-45t45-19q40 0 224 128q32 21 53 35.5t57 45.5t57.5 59t39 68t17.5 80z"/>`),
		g.Group(children),
	)
}
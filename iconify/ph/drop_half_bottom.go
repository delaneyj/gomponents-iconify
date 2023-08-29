package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropHalfBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174 47.75a254.19 254.19 0 0 0-41.45-38.3a8 8 0 0 0-9.18 0A254.19 254.19 0 0 0 82 47.75C54.51 79.32 40 112.6 40 144a88 88 0 0 0 176 0c0-31.4-14.51-64.68-42-96.25ZM187.83 184H68.17a72 72 0 0 1-8-16h135.7a72 72 0 0 1-8.04 16ZM200 144a70.57 70.57 0 0 1-.46 8H56.46a70.57 70.57 0 0 1-.46-8q0-4 .36-8h143.28q.36 4 .36 8ZM128 26c14.16 11.1 56.86 47.74 68.84 94H59.16C71.14 73.76 113.84 37.12 128 26ZM82.81 200h90.38a71.82 71.82 0 0 1-90.38 0Z"/>`),
		g.Group(children),
	)
}
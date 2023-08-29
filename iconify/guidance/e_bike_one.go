package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EBikeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M16.5 5.5h-9V6c0 5-3.5 6.5-3.5 6.5m7.5-.5a5 5 0 0 0-5-5a6 6 0 0 0-6 6a9 9 0 0 0 9 9h.5m0 0c0 1 .5 1.5.5 1.5H13M10 22c0-1 .5-1.5.5-1.5H13m0 0v3m0-3h2m-2 3h2M11.498 12a5 5 0 0 1-5 5m0 0a4.978 4.978 0 0 1-3-1m3 1a4.98 4.98 0 0 1-2.098-.46M5 3h4M20 .5h-3.5V6c0 5 3.5 6.5 3.5 6.5M18.5 17a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		g.Group(children),
	)
}
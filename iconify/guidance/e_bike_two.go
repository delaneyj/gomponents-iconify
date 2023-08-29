package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EBikeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M16.5 5.5h-9V6c0 5-3.5 6.5-3.5 6.5m6 9.5h-.5a9 9 0 0 1-9-9a6.003 6.003 0 0 1 4-5.659M10 22c0 1 .5 1.5.5 1.5H13M10 22c0-1 .5-1.5.5-1.5H13m0 0v3m0-3h2m-2 3h2M6.498 17a4.978 4.978 0 0 1-3-1m3 1a4.98 4.98 0 0 1-2.098-.46m2.098.46a5.001 5.001 0 0 0 4.584-3M5 3h4M20 .5h-3.5V6c0 5 3.5 6.5 3.5 6.5M.5.5l13.356 13.356M23.5 23.5l-6.856-6.856m0 0a5 5 0 1 0-2.789-2.789m2.79 2.79l-2.79-2.79"/>`),
		g.Group(children),
	)
}
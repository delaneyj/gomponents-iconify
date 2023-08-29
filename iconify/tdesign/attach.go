package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.314 3.121a5 5 0 1 1 7.07 7.071l-7.777 7.778a3 3 0 1 1-4.243-4.242l7.778-7.778l1.414 1.414l-7.778 7.778a1 1 0 1 0 1.414 1.414l7.779-7.778a3 3 0 1 0-4.243-4.243L4.95 12.314a5 5 0 0 0 7.07 7.07l8.486-8.485l1.414 1.415l-8.485 8.485a7 7 0 0 1-9.9-9.9l7.779-7.778Z"/>`),
		g.Group(children),
	)
}
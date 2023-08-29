package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M5 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM5 6h5m9 9L16 5h-3M9 9l-4 6h7c0-3 2-6 5-6H9Zm0 0L7 6"/>`),
		g.Group(children),
	)
}
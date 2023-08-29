package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barchart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.25 8.375V28h6.5V8.375h-6.5zM12.25 28h6.5V4.125h-6.5V28zm-9 0h6.5V12.625h-6.5V28z"/>`),
		g.Group(children),
	)
}
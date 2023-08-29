package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextExpandSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1-.5-.5Zm0 12a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1-.5-.5Zm7-4a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5Zm0-4a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5ZM4.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7ZM5 6.5a.5.5 0 0 0-1 0V8H2.5a.5.5 0 0 0 0 1H4v1.5a.5.5 0 0 0 1 0V9h1.5a.5.5 0 0 0 0-1H5V6.5Z"/>`),
		g.Group(children),
	)
}
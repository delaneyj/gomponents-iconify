package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 6.5a.5.5 0 0 1 1 0V8a.5.5 0 0 1-1 0V6.5ZM6 3.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5ZM1 6a5 5 0 1 1 10 0A5 5 0 0 1 1 6Zm5-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}
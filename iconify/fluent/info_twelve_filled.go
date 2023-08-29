package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 6A5 5 0 1 1 1 6a5 5 0 0 1 10 0Zm-5.5.5V8a.5.5 0 0 0 1 0V6.5a.5.5 0 0 0-1 0ZM6 3.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z"/>`),
		g.Group(children),
	)
}
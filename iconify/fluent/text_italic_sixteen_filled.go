package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.8 2H7a.75.75 0 0 0 0 1.5h2.01l-3.428 9H3.2a.75.75 0 0 0 0 1.5H9a.75.75 0 0 0 0-1.5H7.188l3.428-9H12.8a.75.75 0 0 0 0-1.5Z"/>`),
		g.Group(children),
	)
}
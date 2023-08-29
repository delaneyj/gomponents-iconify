package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchVisualSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.75A2.75 2.75 0 0 1 4.75 2h.5a.75.75 0 0 1 0 1.5h-.5c-.69 0-1.25.56-1.25 1.25v.5a.75.75 0 0 1-1.5 0v-.5Zm12 6.5A2.75 2.75 0 0 1 11.25 14h-.5a.75.75 0 0 1 0-1.5h.5c.69 0 1.25-.56 1.25-1.25v-.5a.75.75 0 0 1 1.5 0v.5Zm0-6.5A2.75 2.75 0 0 0 11.25 2h-.5a.75.75 0 0 0 0 1.5h.5c.69 0 1.25.56 1.25 1.25v.5a.75.75 0 0 0 1.5 0v-.5ZM4.75 14A2.75 2.75 0 0 1 2 11.25v-.5a.75.75 0 0 1 1.5 0v.5c0 .69.56 1.25 1.25 1.25h.5a.75.75 0 0 1 0 1.5h-.5ZM8 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM5.25 6a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z"/>`),
		g.Group(children),
	)
}
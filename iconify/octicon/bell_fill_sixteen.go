package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellFillSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16c.9 0 1.7-.6 1.9-1.5c.1-.3-.1-.5-.4-.5h-3c-.3 0-.5.2-.4.5c.2.9 1 1.5 1.9 1.5ZM3 5c0-2.8 2.2-5 5-5s5 2.2 5 5v3l1.7 2.6c.2.2.3.5.3.8c0 .8-.7 1.5-1.5 1.5h-11c-.8.1-1.5-.6-1.5-1.4c0-.3.1-.6.3-.8L3 8.1V5Z"/>`),
		g.Group(children),
	)
}
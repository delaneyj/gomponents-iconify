package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 1.5V2h1v-.5H7Zm1-.01v-.5H7v.5h1ZM8 13.5V4H7v9.5h1Zm0-12v-.01H7v.01h1ZM4 5h3.5V4H4v1Zm-2 9h11v-1H2v1Z"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 4.5V5h1v-.5H7Zm1-.01v-.5H7v.5h1ZM8 11V7H7v4h1Zm0-6.5v-.01H7v.01h1ZM6 8h1.5V7H6v1Zm0 3h3v-1H6v1Z"/>`),
		g.Group(children),
	)
}
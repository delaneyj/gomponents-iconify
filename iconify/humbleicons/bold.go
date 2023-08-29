package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11.5V4h6.25A3.75 3.75 0 0 1 17 7.75v0a3.75 3.75 0 0 1-3.75 3.75H7Zm0 0V20h6.75A4.25 4.25 0 0 0 18 15.75v0a4.25 4.25 0 0 0-4.25-4.25H7Z"/>`),
		g.Group(children),
	)
}
package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.006 20.005l.01-.01m-4.01.01l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m3.99.01l.01-.01m3.99 16.01h8v-16h-8v16Z"/>`),
		g.Group(children),
	)
}
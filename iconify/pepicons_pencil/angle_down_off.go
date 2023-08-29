package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDownOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="m15.68 7.116l-6 5l.64.768l6-5l-.64-.768Z"/><path d="m16.32 7.884l-6 5c-.512.427-1.152-.341-.64-.768l6-5c.512-.427 1.152.341.64.768Z"/><path d="m3.68 7.884l6 5l.64-.768l-6-5l-.64.768Z"/><path d="m4.32 7.116l6 5c.512.427-.128 1.195-.64.768l-6-5c-.512-.427.128-1.195.64-.768ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
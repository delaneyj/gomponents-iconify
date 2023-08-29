package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="m16.32 12.116l-6-5l-.64.768l6 5l.64-.768Z"/><path d="m15.68 12.884l-6-5c-.512-.427.128-1.195.64-.768l6 5c.512.427-.128 1.195-.64.768Z"/><path d="m4.32 12.884l6-5l-.64-.768l-6 5l.64.768Z"/><path d="m3.68 12.116l6-5c.512-.427 1.152.341.64.768l-6 5c-.512.427-1.152-.341-.64-.768Z"/></g>`),
		g.Group(children),
	)
}
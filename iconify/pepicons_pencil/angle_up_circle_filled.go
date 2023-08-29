package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleUpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilAngleUpCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="m19.32 15.116l-6-5l-.64.768l6 5l.64-.768Z"/><path d="m18.68 15.884l-6-5c-.512-.427.128-1.195.64-.768l6 5c.512.427-.128 1.195-.64.768Z"/><path d="m7.32 15.884l6-5l-.64-.768l-6 5l.64.768Z"/><path d="m6.68 15.116l6-5c.512-.427 1.152.341.64.768l-6 5c-.512.427-1.152-.341-.64-.768Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilAngleUpCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
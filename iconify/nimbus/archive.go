package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.5 2.75a1.25 1.25 0 0 0-1.25-1.25H1.75A1.25 1.25 0 0 0 .5 2.75v2.5h1.25v8A1.25 1.25 0 0 0 3 14.5h10a1.25 1.25 0 0 0 1.25-1.25v-8h1.25zM13 5.25v8H3v-8h10zM14.25 4H1.75V2.75h12.5z"/><path fill="currentColor" d="M7.38 8.38h1.24a.63.63 0 0 0 0-1.26H7.38a.63.63 0 0 0 0 1.26z"/>`),
		g.Group(children),
	)
}
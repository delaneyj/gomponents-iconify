package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.21.093a.5.5 0 0 1 .58 0l7 5A.5.5 0 0 1 15 5.5v9a.5.5 0 0 1-.5.5H13V7H2v8H.5a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .21-.407l7-5Z"/><path fill="currentColor" fill-rule="evenodd" d="M3 15h9v-4H3v4Zm6-2H6v-1h3v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M12 10V8H3v2h9Z"/>`),
		g.Group(children),
	)
}
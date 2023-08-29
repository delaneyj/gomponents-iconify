package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nmc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#186C9D"/><path fill="#FFF" fill-rule="nonzero" d="m19.261 23.5l.001-.002a1.8 1.8 0 0 0 .458-.05c.876-.205 1.617-.97 1.793-1.796L25 8.556l-2.772-.014l-2.286 8.568l-6.18-8.597l-.004.004l.003-.01L12.74 8.5v.001a1.9 1.9 0 0 0-.459.049c-.875.206-1.616.971-1.793 1.796L7 23.445l2.773.012l2.285-8.568l6.18 8.598h.003l1.02.013zm-6.593-10.894l.483-1.81l6.181 8.599l-.483 1.81l-6.18-8.6z"/></g>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EslintSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 8.732V6.268L7.5 4.6L10 6.268v2.464L7.5 10.4L5 8.732Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.053 1.276A.5.5 0 0 1 3.5 1h8a.5.5 0 0 1 .447.276l3 6a.5.5 0 0 1 0 .448l-3 6A.5.5 0 0 1 11.5 14h-8a.5.5 0 0 1-.447-.276l-3-6a.5.5 0 0 1 0-.448l3-6ZM11 5.732L7.5 3.4L4 5.732v3.536L7.5 11.6L11 9.268V5.732Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Library(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 206q81-75 192-75v234q-110 0-192 76q-81-76-192-76V131q111 0 192 75zm0-75q-27 0-45.5-19T128 66.5t18.5-45T192 3t45.5 18.5t18.5 45t-18.5 45.5t-45.5 19z"/>`),
		g.Group(children),
	)
}
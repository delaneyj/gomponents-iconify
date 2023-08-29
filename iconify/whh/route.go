package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Route(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.665 0zm-512 0zm256 64q106 0 181-18.5t75-45.5v256H.665V0q0 27 75 45.5t181 18.5t181-18.5t75-45.5q0 27 75 45.5t181 18.5zm-256 960q-94-27-168.5-65.5t-125-78.5t-88.5-97t-59.5-104t-36.5-116.5t-21.5-118T.665 320h1024q-6 76-12.5 124.5t-21.5 118t-36.5 116.5t-59.5 104t-88.5 97t-125 78.5t-168.5 65.5z"/>`),
		g.Group(children),
	)
}
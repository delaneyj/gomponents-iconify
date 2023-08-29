package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.03 4.83a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 0 0 1.06 1.06l2.08-2.08h11.413a4.478 4.478 0 1 0-.19-8.951a5.38 5.38 0 0 0-.437-1.834L21.03 4.83Zm-8.098-.122c1.107 0 2.136.333 2.993.903a.24.24 0 0 1 .032.371L4.097 17.843a.243.243 0 0 1-.3.038A5.875 5.875 0 0 1 8.38 7.195a5.405 5.405 0 0 1 4.552-2.487Z"/>`),
		g.Group(children),
	)
}
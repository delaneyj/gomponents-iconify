package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 3H2C.899 3 0 3.9 0 5v10c0 1.1.899 2 2 2h16c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 12H2V9h16v6zm0-9H2V5h16v1zM4 11.1v.6h.6v-.6H4zm3.6 1.199v.601h1.2v-.601h.6v-.6h.6v-.6H8.8v.6h-.601v.6H7.6zm2.4.601v-.601h-.601v.601H10zm-3 0v-.601H5.8v.601H7zm.6-1.201h.6v-.6H7v1.199h.6v-.599zm-2.401.6H5.8v-.6h.6v-.6H5.2v.6h-.6v.6H4v.601h1.199v-.601z"/>`),
		g.Group(children),
	)
}
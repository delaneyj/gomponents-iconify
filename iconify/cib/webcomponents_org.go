package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcomponentsOrg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15.641 2.969l-.016.021H7.49L0 15.974l7.49 12.979h17.063l.057.078l2.844-4.979l.693-1.188h-.01l.01-.021l-.839-1h1.198l.005-.016l.01.016l3.474-5.854l-3.474-5.849l-.005.005l-.01-.026h-1.26l.839-1.021l-3.474-6.089l-.01.01l-.031-.052zm-3.5 6.13h8.255l-.708 1.021h-.031l3.198 5.349h1.167l-.37.438l.474.583h-1.271l-3.198 5.354h.036l.693 1.016H12.1l-3.99-6.891l4.031-6.875z"/>`),
		g.Group(children),
	)
}
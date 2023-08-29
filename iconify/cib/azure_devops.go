package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzureDevops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m0 11.865l2.995-3.953l11.208-4.557V.063l9.828 7.188l-20.078 3.896v10.969L0 20.975zm32-5.933v19.536l-7.672 6.531l-12.401-4.073v4.073l-7.974-9.885l20.078 2.396V7.25z"/>`),
		g.Group(children),
	)
}
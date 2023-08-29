package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZohoInvoice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.762 42.335h21.215c5.812 0 10.523-4.712 10.523-10.523V5.665H21.292a5.262 5.262 0 0 0-5.262 5.262l-.007 5.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.762 16.027h5.261v21.046a5.264 5.264 0 0 1-5.261 5.262h0A5.264 5.264 0 0 1 5.5 37.073V21.288a5.264 5.264 0 0 1 5.262-5.261ZM5.5 25.391h10.523m6.414 0h13.361m-13.361 6.786h13.361M22.437 13.315h5.838v5.838h-5.838z"/>`),
		g.Group(children),
	)
}
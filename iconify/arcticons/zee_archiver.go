package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeeArchiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.723 42.5L8.648 31.457L6.326 12.931l14.397 11.095V42.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.326 12.931L28.412 5.5l13.262 8.721l-1.496 16.358L20.723 42.5m0-18.474l20.951-9.805"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.326 12.931c2.528 1.238 8.042 3.388 12.687 2.253c-3.406 2.683.636 7.603 1.71 8.842m5.051-11.579l8.498-3.093"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.628 15.765c-.02 0-1.374-1.315-1.374-1.315l2.07-.755l1.22 1.155l-1.916.916Zm1.058-1.89s3.555-2.628 4.915-1.312s-4.281 2.032-4.281 2.032"/>`),
		g.Group(children),
	)
}
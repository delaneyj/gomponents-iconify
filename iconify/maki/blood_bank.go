package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11.2 6.1L7.5 1L3.8 6.1c-.5.7-.8 1.6-.8 2.5C3 11 5 13 7.5 13S12 11 12 8.6c0-.9-.3-1.8-.8-2.5zM10 9H8v2H7V9H5V8h2V6h1v2h2z"/>`),
		g.Group(children),
	)
}
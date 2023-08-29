package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="M39.938 46.844L57.95 34.393l5.657 5.657l-12.513 18.013"/><path fill="#D0CFCE" d="M51.538 60C49.718 51.987 42.564 46 34 46s-15.717 5.987-17.538 14h35.076z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2"><path d="M10 60h48m-6.462 0C49.718 51.987 42.564 46 34 46s-15.717 5.987-17.538 14h35.076z"/><path stroke-linejoin="round" d="M39.938 46.844L57.95 34.393l5.657 5.657l-12.513 18.013"/></g>`),
		g.Group(children),
	)
}
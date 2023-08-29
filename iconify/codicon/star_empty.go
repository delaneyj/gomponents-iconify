package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.595 6.252L8 1L6.405 6.252H1l4.373 3.4L3.75 15L8 11.695L12.25 15l-1.623-5.348L15 6.252H9.595zm-7.247.47H6.72L8 2.507L6.72 6.722H2.348zm3.537 2.75l-1.307 4.305l1.307-4.305zm7.767-2.75H9.28h4.372zm-8.75.9h2.366L8 5.214l.732 2.41h2.367l-1.915 1.49l.731 2.409L8 10.032l-1.915 1.49l.731-2.41l-1.915-1.49z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
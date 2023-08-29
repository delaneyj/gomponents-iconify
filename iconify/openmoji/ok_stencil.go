package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OkStencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#B1CC33" d="M59 60H13a1.002 1.002 0 0 1-1-1V13a1.002 1.002 0 0 1 1-1h46c.55 0 .998.447 1 1v46a.945.945 0 0 1-1 1z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" d="M59 60H13a1.002 1.002 0 0 1-1-1V13a1.002 1.002 0 0 1 1-1h46c.55 0 .998.447 1 1v46a.945.945 0 0 1-1 1z"/><path d="M39.099 41.6a1 1 0 0 1-1-1v-9.4a1 1 0 1 1 2 0v9.4a1 1 0 0 1-.999 1h-.001zm7.224-1.645l-3.714-4.41l3.205-3.69a1 1 0 0 0-1.51-1.312l-2.912 3.353v3.307l3.402 4.041a1 1 0 1 0 1.53-1.289zm-16.46-.451a2.363 2.363 0 0 1-1.754-2.305v-2.7a2.362 2.362 0 0 1 1.754-2.304v-2.032a4.333 4.333 0 0 0-3.753 4.336v2.7a4.333 4.333 0 0 0 3.753 4.337v-2.032zm1.292-9.341v2.032A2.362 2.362 0 0 1 32.91 34.5v2.7a2.363 2.363 0 0 1-1.753 2.305v2.032A4.333 4.333 0 0 0 34.91 37.2v-2.7a4.333 4.333 0 0 0-3.753-4.336z"/>`),
		g.Group(children),
	)
}
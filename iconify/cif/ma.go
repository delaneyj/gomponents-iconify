package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#C1272D" d="M.5.5h300v200H.5z"/><path fill="#006233" d="m150.49 67.166l-7.479 23.094l-24.26.021l19.615 14.271l-7.479 23.115a5499.298 5499.298 0 0 1 19.615-14.271l19.635 14.26l-7.5-23.094l19.615-14.292h-24.24l-7.522-23.104zm0 15.813l2.385 7.281l-4.75.01l2.365-7.291zm-9.073 12.167l-1.458 4.531l-6.198-4.51l7.656-.021zm5.125 0h7.917l2.448 7.542l-6.406 4.667l-6.406-4.667l2.447-7.542zm13.042 0h7.646l-6.177 4.531l-1.469-4.531zm-17.073 12.437l3.854 2.781l-6.198 4.51l2.344-7.291zm16 0l2.344 7.292l-6.188-4.5v-.01l3.844-2.782z"/></g>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhatshotOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.65 0-3.188-.513T6 20l1.45-1.45q1.05.725 2.2 1.088T12 20q3.325 0 5.663-2.337T20 12q0-3.325-2.337-5.663T12 4Q8.675 4 6.337 6.337T4 12H2q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.888.788t3.174 2.15q1.363 1.362 2.15 3.175T22 12q0 2.05-.788 3.875t-2.15 3.188q-1.362 1.362-3.175 2.15T12 22Zm-8.025-4.075L8.05 13.85l3 2.5L16 11.4V14h2V8h-6v2h2.6l-3.65 3.65l-3-2.5l-5.025 5.025q.275.575.488.938t.562.812ZM12 12Z"/>`),
		g.Group(children),
	)
}
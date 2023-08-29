package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttributionOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8.5q-1.325 0-2.038.363T9.25 9.9V14q0 .2.15.35t.35.15h1v3.25q0 .525.363.888T12 19q.525 0 .888-.363t.362-.887V14.5h1q.2 0 .35-.15t.15-.35V9.9q0-.675-.713-1.038T12 8.5ZM12 22q-2.05 0-3.875-.788t-3.188-2.15q-1.362-1.362-2.15-3.187T2 12q0-2.075.788-3.888t2.15-3.174Q6.3 3.575 8.124 2.788T12 2q2.075 0 3.888.788t3.174 2.15q1.363 1.362 2.15 3.175T22 12q0 2.05-.788 3.875t-2.15 3.188q-1.362 1.362-3.175 2.15T12 22Zm0-2q3.325 0 5.663-2.337T20 12q0-3.325-2.337-5.663T12 4Q8.675 4 6.337 6.337T4 12q0 3.325 2.337 5.663T12 20Zm0-12q.65 0 1.075-.425T13.5 6.5q0-.65-.425-1.075T12 5q-.65 0-1.075.425T10.5 6.5q0 .65.425 1.075T12 8Zm0 4Z"/>`),
		g.Group(children),
	)
}
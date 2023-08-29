package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.2 8.375l3.5-6q2.275.6 4.038 2.2t2.562 3.8H11.2Zm-2.775 2.5L5 4.875q1.35-1.325 3.138-2.1T12 2q.325 0 .75.038t.775.087l-5.1 8.75Zm-6.1 3.625q-.15-.6-.238-1.225T2 12q0-1.775.575-3.35T4.2 5.775L9.25 14.5H2.325Zm7 7.125q-2.275-.6-4.05-2.2t-2.575-3.8h10.075l-3.45 6ZM12 22q-.375 0-.763-.05t-.737-.1l5.075-8.725l3.425 6q-1.35 1.325-3.138 2.1T12 22Zm7.8-3.775L14.75 9.5h6.925q.15.6.238 1.225T22 12q0 1.75-.613 3.35T19.8 18.225Z"/>`),
		g.Group(children),
	)
}
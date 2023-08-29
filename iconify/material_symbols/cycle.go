package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.85 19.125q-2.6-1.2-4.237-3.625t-1.638-5.475q0-.65.075-1.275t.225-1.225l-1.15.675l-1-1.725L4.9 3.725l2.75 4.75l-1.75 1l-1.35-2.35q-.275.675-.412 1.4T4 10.025q0 2.425 1.325 4.413t3.525 2.937l-1 1.75ZM15.5 7V5h2.725q-1.15-1.425-2.775-2.212T12 2q-1.375 0-2.6.425t-2.25 1.2l-1-1.75Q7.4 1 8.863.5t3.112-.5q1.975 0 3.788.738T19 2.875V1.5h2V7h-5.5Zm-.65 15l-4.775-2.75l2.75-4.75l1.725 1l-1.425 2.45q2.95-.425 4.913-2.663T20 10.026q0-.275-.013-.525t-.062-.5h2.025q.025.25.038.488T22 10q0 3.375-2.013 6.038t-5.237 3.587l1.1.65l-1 1.725Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApprovalDelegation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.925 11.05L8.675 6.8L10.1 5.4l2.825 2.825l5.675-5.65l1.4 1.4l-7.075 7.075ZM14 22.5l-7-1.95V11h1.95l6.2 2.3q.825.3 1.337 1.05T17 16h-2q-1.05 0-1.65-.075T12.3 15.7l-2-.65l-.3.95l1.575.575q.7.275 1.275.35T14.2 17H19q1.65 0 2.325.537T22 19v1l-8 2.5ZM1 22V11h4v11H1Z"/>`),
		g.Group(children),
	)
}
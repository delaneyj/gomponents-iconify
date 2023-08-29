package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DraftOrdersRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.975 11.125l-2.1-2.1l-5.15 5.15q-.15.15-.225.338t-.075.387v1.175q0 .2.15.35t.35.15H9.1q.2 0 .388-.075t.337-.225l5.15-5.15Zm.6-.6l.7-.7q.3-.3.3-.7t-.3-.7l-.7-.7q-.3-.3-.7-.3t-.7.3l-.7.7l2.1 2.1ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}
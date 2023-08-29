package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayDisabledOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.45 13.6l-1.425-1.45L8 5.15V5l9.675 6.15q.475.275.462.838t-.487.862l-1.2.75Zm-6.9 4.425q-.5.325-1.025.038T8 17.175V10.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.3.3.288.7t-.313.7q-.3.275-.713.263t-.687-.288L13 15.8l-3.45 2.225ZM10 12.8Zm0 2.55l1.55-1L10 12.8v2.55Zm5.025-3.2Z"/>`),
		g.Group(children),
	)
}
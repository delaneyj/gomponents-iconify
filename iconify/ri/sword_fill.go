package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwordFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.048 13.406l3.535 3.536l-1.413 1.414l1.415 1.415l-1.414 1.414l-2.475-2.475l-2.829 2.829l-1.414-1.414l2.829-2.83l-2.475-2.474l1.414-1.414l1.414 1.413l1.413-1.414ZM3 3l3.546.003l11.817 11.818l1.415-1.414l1.415 1.414l-2.475 2.475l2.828 2.829l-1.414 1.414l-2.829-2.829l-2.474 2.475l-1.415-1.414l1.414-1.415L3.002 6.531L2.999 3Zm14.457 0L21 3.003l.002 3.523l-4.053 4.052l-3.536-3.535L17.456 3Z"/>`),
		g.Group(children),
	)
}
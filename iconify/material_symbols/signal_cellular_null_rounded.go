package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularNullRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.825 20h13.15V6.85L6.825 20Zm-2.4 2q-.675 0-.938-.613T3.7 20.3L20.275 3.725q.475-.475 1.088-.212t.612.937V21q0 .425-.288.713t-.712.287H4.425Z"/>`),
		g.Group(children),
	)
}
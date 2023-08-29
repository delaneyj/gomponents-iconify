package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMoreRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18.1l2.325-2.325q.3-.3.725-.3t.725.3q.3.3.3.725t-.3.725L12.7 20.3q-.15.15-.325.213t-.375.062q-.2 0-.375-.063T11.3 20.3l-3.075-3.075q-.3-.3-.3-.725t.3-.725q.3-.3.725-.3t.725.3L12 18.1ZM12 6L9.675 8.325q-.3.3-.725.3t-.725-.3q-.3-.3-.3-.725t.3-.725L11.3 3.8q.15-.15.325-.213T12 3.526q.2 0 .375.063t.325.212l3.075 3.075q.3.3.3.725t-.3.725q-.3.3-.725.3t-.725-.3L12 6Z"/>`),
		g.Group(children),
	)
}
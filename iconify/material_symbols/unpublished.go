package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unpublished(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-2.95-2.95q-1.2.8-2.587 1.225T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.55.425-2.938T3.65 6.476L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425Zm-.1-5.8l-5.325-5.35l2.6-2.6l-1.4-1.4l-2.6 2.625l-7.15-7.15q1.2-.775 2.588-1.2T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.525-.425 2.913t-1.2 2.587Zm-9.775-.9l1.6-1.6l-1.4-1.4l-.2.2l-2.85-2.85l-1.4 1.4l4.25 4.25Z"/>`),
		g.Group(children),
	)
}
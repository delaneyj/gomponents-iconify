package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlinesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.8 18h11.55L19.6 6h-5.55L5.8 18Zm-3.25 1.225l9.85-14.35q.275-.425.713-.65T14.05 4h5.55q.95 0 1.538.725t.412 1.65l-2.4 12.8q-.075.35-.35.588t-.625.237H2.95q-.3 0-.438-.263t.038-.512ZM14.5 14q1.05 0 1.775-.725T17 11.5q0-1.05-.725-1.775T14.5 9q-1.05 0-1.775.725T12 11.5q0 1.05.725 1.775T14.5 14Zm-8.7 4h11.55H5.8Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndArrowNotchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.7 17.925q-.35.2-.625-.063T12 17.25L14.425 13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h11.425L12 6.75q-.2-.35.075-.613t.625-.062l7.975 5.075q.475.3.475.85t-.475.85L12.7 17.925Z"/>`),
		g.Group(children),
	)
}
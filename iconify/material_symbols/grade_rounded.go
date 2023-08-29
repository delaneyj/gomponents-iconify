package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.05 21.05q-.275.225-.587.025t-.188-.55L8.15 14.4l-4.875-3.5q-.3-.2-.187-.55T3.55 10H9.6l1.925-6.4q.05-.2.188-.275T12 3.25q.15 0 .288.075t.187.275L14.4 10h6.05q.35 0 .463.35t-.188.55l-4.875 3.5l1.875 6.125q.125.35-.188.55t-.587-.025L12 17.3l-4.95 3.75Z"/>`),
		g.Group(children),
	)
}
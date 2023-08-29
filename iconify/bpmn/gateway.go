package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gateway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<rect width="12.824" height="12.824" x="-736.344" y="737.679" fill="transparent" stroke="currentColor" stroke-width="1.15" ry=".199" transform="rotate(-45 -120466.552 -49977.13) scale(96.7529)"/>`),
		g.Group(children),
	)
}
package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="20" fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="36" cy="36" r="7"/><circle cx="36" cy="36" r="7" fill="#D22F27"/>`),
		g.Group(children),
	)
}
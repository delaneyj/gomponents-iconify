package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.604 2.76a20.5 20.5 0 0 0-4.637 0l-1.595.182a.5.5 0 0 0-.441.453l-.123 1.382a39.5 39.5 0 0 0 0 6.983l.123 1.382a.5.5 0 0 0 .498.456H10.5V21a.5.5 0 0 0 .89.312l.391-.49a35.5 35.5 0 0 0 5.497-9.676l.19-.507A.5.5 0 0 0 17 9.963h-2.713l2.325-6.352a.5.5 0 0 0-.413-.669l-1.595-.181Z"/>`),
		g.Group(children),
	)
}
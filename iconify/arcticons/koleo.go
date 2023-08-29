package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Koleo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M16.8 43.5a5 5 0 0 1-3.536-8.535l10.931-10.932l-10.997-10.998a5 5 0 0 1 7.071-7.07l14.533 14.533a5 5 0 0 1 0 7.07L20.335 42.035a4.985 4.985 0 0 1-3.536 1.465Z"/>`),
		g.Group(children),
	)
}
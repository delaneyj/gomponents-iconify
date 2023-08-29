package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taigamobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 28.378h9.048l-5.926-10.263"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.115h5.872L24 7.944v31.48h13.525l-6.378-11.046M24 24.934h-9.048l5.926-10.264"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.67h-5.872L24 4.5v31.48H10.475l6.378-11.046"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.988 35.98v7.52h6.024v-4.076"/>`),
		g.Group(children),
	)
}
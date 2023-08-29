package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mitidapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="14.185" cy="17.762" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.112" ry="5.159"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.185 26.048c-5.315.027-9.707 3.239-9.685 9.349h19.37c.015-5.947-4.37-9.376-9.685-9.349Zm12.668-13.445v22.794h6.514C40.077 35.415 43.53 30.254 43.5 24s-3.24-11.422-10.133-11.396h-6.514Z"/>`),
		g.Group(children),
	)
}
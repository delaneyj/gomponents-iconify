package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screentime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="15.695" height="27.197" x="16.152" y="10.401" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.881"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.07 7.623L24 24l20.92-4.96"/>`),
		g.Group(children),
	)
}
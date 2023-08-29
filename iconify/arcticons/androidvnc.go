package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Androidvnc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.7 30.13V17.87l8.12 12.26V17.87m-11.19 0l-4.07 12.26L8.5 17.87m31 8.13v.05a4.06 4.06 0 0 1-4.06 4.06h0a4.07 4.07 0 0 1-4.07-4.06v-4.12a4.07 4.07 0 0 1 4.07-4.06h0a4.06 4.06 0 0 1 4.06 4.06V22"/>`),
		g.Group(children),
	)
}
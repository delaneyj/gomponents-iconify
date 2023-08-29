package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouterKeygen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="12.902" cy="36.332" r="3.029" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.126 42.11v-5.778h0h-22.195m18.404 3.331v-3.331m-3.917-9.599a9.976 9.976 0 0 0-12.746-.035v.035m18.83-6.916a19.615 19.615 0 0 0-24.965 0m30.963-7.26a29 29 0 0 0-37 0"/><circle cx="24" cy="31.441" r="1.423" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
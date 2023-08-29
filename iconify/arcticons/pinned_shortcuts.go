package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinnedShortcuts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.979 4.5h19.797v4.092h-2.029v15.32l4.058 3.988v3.917H25.522V43.5h-3.253V31.818H12.195v-4.023l3.813-3.813V8.593h-2.029V4.5Z"/>`),
		g.Group(children),
	)
}
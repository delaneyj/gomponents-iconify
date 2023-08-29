package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Recon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.5 5.5h-21a8 8 0 0 0-8 8v21a8 8 0 0 0 8 8h21a8 8 0 0 0 8-8v-21a8 8 0 0 0-8-8Z"/><rect width="15" height="15" x="16.5" y="16.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}
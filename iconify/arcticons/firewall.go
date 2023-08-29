package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firewall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.1 14.5v-8a2 2 0 0 0-2-2H14.9a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h18.2a2 2 0 0 0 2-2v-8m-22.2-24h22.2m-22.2 29h22.2"/><circle cx="35.1" cy="24" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.5 28.6l9.2-9.2"/>`),
		g.Group(children),
	)
}
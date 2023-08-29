package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThunderVpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.979 43.5l16.874-21.522H23.805l4.482-9.244h9.548L41.944 4.5H10.165l-4.109 8.234h12.983c-2.485 5.004-4.977 10.004-7.555 14.968l9.888-.067L16.98 43.5"/>`),
		g.Group(children),
	)
}
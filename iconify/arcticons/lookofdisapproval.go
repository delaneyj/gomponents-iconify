package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lookofdisapproval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="M18.934 32.448h10.132"/><circle cx="12.581" cy="25.151" r="5.617" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M6.534 19.534h9.958A1.91 1.91 0 0 0 17.51 16"/><circle cx="35.647" cy="25.151" r="5.617" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M29.6 19.534h9.958A1.91 1.91 0 0 0 40.576 16"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.581" cy="25.151" r=".75" fill="currentColor"/><circle cx="35.647" cy="25.151" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}
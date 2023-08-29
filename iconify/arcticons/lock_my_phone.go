package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockMyPhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.648 41.49V22.368h3.992a.858.858 0 0 0 .556-1.512L25.59 5.868a2.452 2.452 0 0 0-3.178 0L4.804 20.855a.858.858 0 0 0 .556 1.512h3.992v19.124a1.226 1.226 0 0 0 1.226 1.226h26.843a1.226 1.226 0 0 0 1.227-1.226Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.732 24.867v-3.47a5.269 5.269 0 0 1 10.537 0v3.47m-12.375 0h14.211v11.675H16.894z"/>`),
		g.Group(children),
	)
}
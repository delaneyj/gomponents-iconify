package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loklok(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 20.546v6.908h3.454m7.517-6.908v6.908m3.713 0L20.34 24l2.844-3.431M20.34 24h-.869"/><rect width="4.577" height="6.908" x="13.033" y="20.546" rx="2.288" ry="2.288"/><path d="M24.816 20.546v6.908h3.454m7.517-6.908v6.908m3.713 0L36.656 24l2.844-3.431M36.656 24h-.869"/><rect width="4.577" height="6.908" x="29.349" y="20.546" rx="2.288" ry="2.288"/></g>`),
		g.Group(children),
	)
}
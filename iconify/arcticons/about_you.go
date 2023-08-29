package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AboutYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M34.674 15.554H39.5m-2.413 7.284v-7.284m-23.761 9.608l-2.413 3.642L8.5 25.162m2.413 7.284v-3.642m17.255-13.25v4.871a2.413 2.413 0 0 0 4.826 0v-4.871m-14.988 3.642a1.821 1.821 0 0 1 0 3.642h-3.005v-7.284h3.005a1.821 1.821 0 0 1 0 3.642h0Zm0 0h-3.005m-2.513 1.205H9.332m-.786 2.416l2.367-7.263l2.368 7.284"/><rect width="4.826" height="7.285" x="21.517" y="15.554" rx="2.413" ry="2.413"/><path d="M21.026 25.162v4.871a2.413 2.413 0 0 0 4.826 0v-4.871"/><rect width="4.826" height="7.285" x="14.375" y="25.162" rx="2.413" ry="2.413"/></g>`),
		g.Group(children),
	)
}
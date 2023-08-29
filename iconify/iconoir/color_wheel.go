package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorWheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Z"/><path d="M12 16a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-14v6m0 8v6M2 12h6m8 0h6M4.929 4.929L9.172 9.17m5.656 5.659l4.243 4.242m-14.142 0l4.243-4.242m5.656-5.658l4.243-4.242"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroadcastOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.364 19.364A9 9 0 0 0 8.643 4.647M6.155 6.156a9 9 0 0 0-.519 13.208"/><path d="M15.536 16.536A5 5 0 0 0 12 8M9 9a5 5 0 0 0-.535 7.536M12 12a1 1 0 1 0 1 1M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
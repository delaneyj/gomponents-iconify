package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnimeSlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.489 41.541C12.07 33.858 20.805 13.639 26.386 4.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.65 43.5c-3.8-6.47-10.11-28.766-10.981-36.026"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.027 30.936c8.411-.151 23.94-3.63 31.946-7.592"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EpicGames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-29a2 2 0 0 0-2 2v27.747a2 2 0 0 0 1.127 1.8L24 44.5l15.373-7.453a2 2 0 0 0 1.127-1.8V7.5a2 2 0 0 0-2-2ZM17.925 38.298h12.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.512 28h-4V12h4m19 5v-3a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2v-3m-16 5V12h3a2 2 0 0 1 2 2v3.967a2 2 0 0 1-2.032 2l-2.968-.065M12.488 20h4m11.024-8v16"/>`),
		g.Group(children),
	)
}
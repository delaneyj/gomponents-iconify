package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tmshappinest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.3 28.69V19.3l4.7 9.4l4.7-9.38v9.38m-18.2-9.4h6.22m-3.11 9.4v-9.4m17.85 8.37a2.64 2.64 0 0 0 2.31 1h1.39a2.35 2.35 0 0 0 2.34-2.35h0A2.35 2.35 0 0 0 35.16 24h-1.54a2.35 2.35 0 0 1-2.34-2.35h0a2.35 2.35 0 0 1 2.34-2.35H35a2.61 2.61 0 0 1 2.3 1"/>`),
		g.Group(children),
	)
}
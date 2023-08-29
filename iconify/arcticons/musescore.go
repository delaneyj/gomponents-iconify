package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musescore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="33" cy="20.4" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 25.35v7.2m0-7.2a4.513 4.513 0 0 1 4.5-4.5h0a4.513 4.513 0 0 1 4.5 4.5m0 0v2.925a4.513 4.513 0 0 0 4.5 4.5h.001a4.512 4.512 0 0 0 4.499-4.5V23.1m-18 2.25a4.513 4.513 0 0 0-4.5-4.5h0a4.513 4.513 0 0 0-4.5 4.5v7.2m18-12.825a4.513 4.513 0 0 1 4.5-4.5h.001a4.512 4.512 0 0 1 4.5 4.5"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessKalender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 21.3v-3.24a3.24 3.24 0 0 0-3.24-3.24h-4.32a3.24 3.24 0 0 0-3.24 3.24v11.88a3.24 3.24 0 0 0 3.24 3.24h4.32a3.24 3.24 0 0 0 3.24-3.24v-4.32m-27 1.08v3.24a3.24 3.24 0 0 0 3.24 3.24h5.4a3.24 3.24 0 0 0 3.24-3.24V26.7a3.24 3.24 0 0 0-3.24-3.24m-8.64-5.4v-3.24h8.64a3.24 3.24 0 0 1 3.24 3.24v2.16a3.24 3.24 0 0 1-3.24 3.24H15.9"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DictionaryCsFr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Zm4.33-2v39H37.6a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Zm7.327 23.482h4m-4 4h2.6m-2.6-4v8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.018 35.982v-8h2.6a2.7 2.7 0 0 1 0 5.4h-2.6m2.789-.007l2.511 2.607m-5.2-16.764a2.384 2.384 0 0 0 2 .9h1.2a2.006 2.006 0 0 0 2-2h0a2.006 2.006 0 0 0-2-2h-1.3a2.006 2.006 0 0 1-2-2h0a2.006 2.006 0 0 1 2-2h1.2a2.147 2.147 0 0 1 2 .9m-6.761 4.4h0a2.689 2.689 0 0 1-2.7 2.7h0a2.689 2.689 0 0 1-2.7-2.7v-2.7a2.689 2.689 0 0 1 2.7-2.7h0a2.606 2.606 0 0 1 2.6 2.7h0"/>`),
		g.Group(children),
	)
}
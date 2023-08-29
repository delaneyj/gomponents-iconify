package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoctailCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M9 11a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1H9Z"/><path fill-rule="evenodd" d="m5.767 8.68l6.5 7a1 1 0 0 0 1.466 0l6.5-7c.594-.64.14-1.68-.733-1.68h-13c-.873 0-1.327 1.04-.733 1.68Zm11.44.32L13 13.53L8.793 9h8.414Z" clip-rule="evenodd"/><path d="M14 14.5v7h-2v-7h2Z"/><path d="M17.5 23h-9c0-1.475 2.05-2.5 4.5-2.5s4.5 1.025 4.5 2.5ZM15.818 4.728a.75.75 0 0 1 .364-1.456l4 1a.75.75 0 0 1-.364 1.456l-4-1Z"/><path d="M14.211 11.737a.75.75 0 1 1-1.423-.474l2.5-7.5a.75.75 0 1 1 1.424.474l-2.5 7.5Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpineLabel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M3 11v2h5.59l-6.3 6.29l1.42 1.42l6.29-6.3V20h2v-9H3z" fill="currentColor"/><path d="M26 13h-3v-1h-2v1h-3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h3v1h2v-1h3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2zm-8 4v-2h8v2z" fill="currentColor"/><path d="M26 23h-3v-1h-2v1h-3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h3v1h2v-1h3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2zm-8 4v-2h8v2z" fill="currentColor"/><path d="M26 3h-3V2h-2v1h-3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h3v1h2V9h3a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm-8 4V5h8v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}
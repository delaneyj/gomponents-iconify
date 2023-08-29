package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaggageClaim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="6" cy="26" r="2" fill="currentColor"/><path fill="currentColor" d="M28 18h-3v-2a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v2h-3a2.002 2.002 0 0 0-2 2v8a2.002 2.002 0 0 0 2 2h14a2.002 2.002 0 0 0 2-2v-8a2.002 2.002 0 0 0-2-2Zm-9-2h4v2h-4Zm9 12H14v-8h14Z"/><path fill="currentColor" d="M10 6h4v6h2V6h4v6h2V6h4v6h2V6.005A2.005 2.005 0 0 0 25.995 4H4.005A2.005 2.005 0 0 0 2 6.005v13.99A2.005 2.005 0 0 0 4.005 22H10ZM8 20H4V6h4Z"/>`),
		g.Group(children),
	)
}
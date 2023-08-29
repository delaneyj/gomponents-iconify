package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.125 19.375h-8.25a.5.5 0 0 0-.5.5v8.25a.5.5 0 0 0 .5.5h8.25a.5.5 0 0 0 .5-.5h0v-8.25a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}
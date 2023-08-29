package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Goodrx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 18.5h-7.5V11a5.5 5.5 0 0 0-11 0v7.5H11a5.5 5.5 0 0 0 0 11h7.5V37a5.5 5.5 0 0 0 11 0v-7.5H37a5.5 5.5 0 0 0 0-11Zm-7.5 11h-11m0-11v11"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pbz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.94 15.633h30.12M8.94 32.367v-8.201a3.605 3.605 0 0 1 7.21 0v8.2m15.7.001v-8.201a3.605 3.605 0 0 1 7.21 0v8.2m-18.665.001v-8.201a3.605 3.605 0 0 1 7.21 0v8.2"/>`),
		g.Group(children),
	)
}
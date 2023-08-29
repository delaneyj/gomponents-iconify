package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markdownr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 32.597V15.403l8.597 17.194l8.598-17.194v17.194m4.137-3.755v3.756H43v-3.756"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.264 23.262l6.134 6.116l6.492-6.116h-3.684v-6.903H31.59v6.903Z"/>`),
		g.Group(children),
	)
}
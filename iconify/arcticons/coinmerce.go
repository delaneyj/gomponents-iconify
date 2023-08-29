package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coinmerce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.493 22.465v-2.214L24 16.503l-6.493 3.748v7.498L24 31.497l6.493-3.748v-2.214h5.63v5.464l-12.123 7l-12.123-7v-5.464h-5.63v8.715L24 44.5l17.754-10.25v-20.5L24 3.5L6.247 13.75v8.715h5.63v-5.464l12.123-7l12.123 7v5.464h-5.63z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onshape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.175 17.75v12.5L24 36.5l10.825-6.25v-12.5L24 11.5l-10.825 6.25z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.247 13.75v20.5L24 44.5l17.754-10.25v-20.5L24 3.5L6.247 13.75z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.67 21.5v5L24 29l4.33-2.5v-5L24 19l-4.33 2.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.639 19.75v8.5L24 32.5l7.361-4.25v-8.5L24 15.5l-7.361 4.25zm18.186 10.5v8M28.33 26.5V30M24 11.5l6.928-4M13.175 30.25l-6.928-4m20.784-9L24 19m-7.361 5.75l3.031 1.75"/>`),
		g.Group(children),
	)
}
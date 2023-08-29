package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libretrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.588 43.5l16.966-10.25l.236-19.088L24.059 4.5L7.21 13.926v19.088l10.604 6.468l16.967-9.78l.118-12.135l-10.84-6.292L13.1 17.272v12.372l4.737 2.934l11.052-6.41v-5.42l-4.831-2.828l-4.949 2.651l-.083 4.784"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Studierendenwerkheidelberg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41 32.5l2.5-24H14c-3.3 0-6.658 3.718-7 7l-2.5 24H34c3.3 0 6.658-3.718 7-7Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.1 20.6a2.384 2.384 0 0 0 2 .9h1.2a2 2 0 0 0 0-4H14a2 2 0 0 1 0-4h1.2a2.147 2.147 0 0 1 2 .9m5.504-.9v8m5.296-8l2 8l2-8l2 8l2-8m-19 17h5.3m3.4 4v-8H27a4 4 0 0 1 0 8Zm-5.7-21h5.408m-3.108 13v8m-5.3-8v8"/>`),
		g.Group(children),
	)
}
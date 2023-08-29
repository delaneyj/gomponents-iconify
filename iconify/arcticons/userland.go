package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Userland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 3.5l17.754 10.25v20.5L24 44.5L6.246 34.25v-20.5L24 3.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.209 34.565c-1.092-10.158-5.286-22.632-10.459-22.632c-3.7 0-4.878 1.794-6.75 5.103c-1.872-3.309-3.05-5.103-6.75-5.103c-5.173 0-9.366 12.474-10.458 22.632"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 30.325l5.733-4.483a.53.53 0 0 0 .008-.837A9.175 9.175 0 0 0 24 23.25a9.175 9.175 0 0 0-5.741 1.757a.53.53 0 0 0 .008.836Z"/>`),
		g.Group(children),
	)
}
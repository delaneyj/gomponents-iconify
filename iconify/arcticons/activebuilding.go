package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Activebuilding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.46 38.63H43.5v-6.59L24 9.38L4.5 31.96"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 9.38v29.25h5.46m-3.04-21.06l3.01 3.39v3.59l-3.02-3.41l.01-3.57zm5.05 5.89l3.01 3.4v3.58l-3.02-3.41l.01-3.57z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.52 24.71l3 3.39v3.59l-3.02-3.41l.02-3.57zm9.89 4.52l3 3.39v3.59l-3.02-3.41l.02-3.57zm-4.92 1.24l3 3.39v3.58l-3.02-3.41l.02-3.56zm-10.08-12.9L18.4 21v3.59l3-3.41Zm-5.05 5.89l-3 3.39v3.59l3-3.41Zm-4.93 5.77l-3 3.39v3.59l3-3.41Zm9.89-4.52l-3 3.39v3.59l3-3.41Zm-4.97 5.76l-3 3.39v3.58l3-3.41Zm5.03 1.18l-3 3.39v3.59l3-3.42Zm5.14 0l3 3.39v3.58l-3.02-3.41l.02-3.56z"/>`),
		g.Group(children),
	)
}
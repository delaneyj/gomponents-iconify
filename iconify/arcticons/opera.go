package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.922 36.116c-2.376-2.806-3.916-6.954-4.021-11.61v-1.012c.104-4.655 1.644-8.803 4.022-11.61C20.005 7.88 24.591 5.34 29.71 5.34c3.15 0 6.096.961 8.616 2.634A21.412 21.412 0 0 0 24.08 2.503L24 2.5C12.125 2.5 2.5 12.127 2.5 24c0 11.528 9.077 20.938 20.476 21.472a21.421 21.421 0 0 0 15.352-5.445c-2.52 1.671-5.466 2.633-8.616 2.633c-5.12 0-9.705-2.54-12.79-6.545v0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.922 11.885c1.973-2.329 4.522-3.733 7.306-3.733c6.257 0 11.332 7.096 11.332 15.849s-5.073 15.848-11.333 15.848c-2.783 0-5.332-1.404-7.304-3.733c3.084 4.006 7.67 6.545 12.79 6.545c3.15 0 6.095-.961 8.615-2.632C42.73 36.092 45.5 30.369 45.5 24s-2.77-12.089-7.17-16.026c-2.522-1.672-5.468-2.634-8.618-2.634c-5.12 0-9.705 2.54-12.79 6.545Z"/>`),
		g.Group(children),
	)
}
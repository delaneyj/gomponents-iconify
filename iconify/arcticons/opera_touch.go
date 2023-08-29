package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OperaTouch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M32.669 42.373c-.96.19-1.948.288-2.957.288c-5.12 0-9.705-2.54-12.79-6.545v0c-2.376-2.806-3.916-6.954-4.021-11.61v-1.012c.104-4.655 1.644-8.804 4.022-11.61C20.006 7.878 24.592 5.34 29.71 5.34c3.15 0 6.096.961 8.616 2.634A21.412 21.412 0 0 0 24.08 2.502L24 2.501C12.125 2.5 2.5 12.127 2.5 24c0 11.528 9.077 20.938 20.476 21.472a21.405 21.405 0 0 0 10.457-2.15"/><path d="M43.326 33.431A21.41 21.41 0 0 0 45.5 24c0-6.368-2.77-12.089-7.17-16.026c-2.521-1.672-5.468-2.634-8.618-2.634c-5.12 0-9.705 2.54-12.79 6.545c1.973-2.329 4.522-3.733 7.306-3.733c4.478 0 8.35 3.634 10.19 8.91c.732 2.095 1.142 4.45 1.142 6.939c0 8.752-5.073 15.848-11.333 15.848c-2.783 0-5.332-1.404-7.304-3.733c3.084 4.006 7.67 6.545 12.79 6.545c1.008 0 1.996-.099 2.956-.287"/></g><circle cx="38.5" cy="38.499" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.85 34.499h5.3m-2.65 8v-8"/>`),
		g.Group(children),
	)
}
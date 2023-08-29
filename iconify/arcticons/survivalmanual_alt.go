package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurvivalmanualAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.48 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5H8.43a2 2 0 0 0-1.95 1.95Zm17.79 24.513A6.963 6.963 0 1 1 31.233 24a6.963 6.963 0 0 1-6.963 6.963Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.35 20.673l-4.107 1.961a.769.769 0 0 0-.339.34l-1.951 4.107a.185.185 0 0 0 .09.245l.002.001a.175.175 0 0 0 .144 0l4.108-1.961a.853.853 0 0 0 .38-.38l1.961-4.108a.195.195 0 0 0-.113-.236a.175.175 0 0 0-.174.03Z"/><path fill="currentColor" d="M24.27 24.75a.75.75 0 1 1 .75-.75a.76.76 0 0 1-.75.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.81 43.5h24.92a2 2 0 0 0 2-2L37.68 29V6.45a2 2 0 0 0-2-1.95H10.81"/>`),
		g.Group(children),
	)
}
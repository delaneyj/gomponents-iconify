package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bikewale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M45.5 24c0 4-1.09 7.74-3 10.94a21.418 21.418 0 0 1-5.95 6.51c-.97.7-2.15 1.05-3.35 1.05H14.8c-1.2 0-2.38-.35-3.35-1.05a21.396 21.396 0 0 1-5.95-6.51c-1.91-3.2-3-6.94-3-10.94s1.09-7.74 3-10.94c1.83-3.11 4.45-5.73 7.56-7.56c3.2-1.91 6.94-3 10.94-3s7.74 1.09 10.94 3c3.11 1.83 5.73 4.45 7.56 7.56c1.91 3.2 3 6.94 3 10.94Z"/><circle cx="24" cy="24" r="5.66"/><path d="m27.53 19.58l13.39-8.83M27.42 28.5l13.04 9.32"/></g>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EasyDiary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.039 18.75a5.513 5.513 0 0 0-.081 4.61L28.132 40.064a2.461 2.461 0 0 1-2.635.778L8.528 35.5a6.723 6.723 0 0 1-2.829-1.565a5.38 5.38 0 0 1-1.055-2.288a9.73 9.73 0 0 1 .013-3.289c.636-3.084 2.91-2.177 2.91-2.177"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.567 26.182l17.305 5.852a2.446 2.446 0 0 0 2.643-.729l15.756-18.444a.994.994 0 0 0-.54-1.616L24.025 7.112a2.754 2.754 0 0 0-2.69.902L5.511 26.581"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m19.748 15.112l12.625 3.602l3.17-3.573l-13.011-3.112Z"/>`),
		g.Group(children),
	)
}
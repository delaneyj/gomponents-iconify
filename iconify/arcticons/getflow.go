package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Getflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.864 15.954A21.452 21.452 0 1 1 32.487 4.321"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.27 19.192a10.479 10.479 0 1 1-4.453-4.458"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.058 24.032a3.09 3.09 0 1 1-3.09-3.09m0 3.09L39.374 8.626m-6.669 6.594V9.17l6.67-6.67v6.126"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.78 15.295h6.05l6.67-6.67h-6.126"/>`),
		g.Group(children),
	)
}
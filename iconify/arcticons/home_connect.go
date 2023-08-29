package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.302 33.464a21.513 21.513 0 1 0-9.838 9.838a2.002 2.002 0 0 1 1.317-.17l7.659 1.692a2 2 0 0 0 2.384-2.384l-1.691-7.659a2.002 2.002 0 0 1 .17-1.317Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.392 19.38l-6.67-5.616a1.121 1.121 0 0 0-1.444 0l-6.67 5.615a1.682 1.682 0 0 0-.598 1.287v10.713a1.121 1.121 0 0 0 1.12 1.121h13.74a1.121 1.121 0 0 0 1.12-1.121V20.666a1.682 1.682 0 0 0-.598-1.287Z"/>`),
		g.Group(children),
	)
}
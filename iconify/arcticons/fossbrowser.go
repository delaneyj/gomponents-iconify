package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fossbrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.46 21.46 0 0 0 24 2.5Zm6 5.34a17.23 17.23 0 0 1 6.21 28.32v-.7c-.11-1.27 0-2.47-3.41-2.52h-3V27c0-1.87-.68-2.59-2.13-2.59h-13v-4.5H20c.53 0 1.21-.68 1.23-1.74v-4.54h6c2.79 0 2.66-2.86 2.78-5ZM7 21.34l9.84 9.23c.11 3-.37 6.52 4.47 6.49v4A17.24 17.24 0 0 1 6.76 24A16.94 16.94 0 0 1 7 21.34Zm0 0"/>`),
		g.Group(children),
	)
}
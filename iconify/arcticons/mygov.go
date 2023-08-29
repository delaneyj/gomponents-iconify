package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mygov(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.366 24L11.168 6.056C9.956 4.861 7.903 5.72 7.903 7.422v33.156c0 1.703 2.053 2.561 3.265 1.366L29.366 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.097 24L21.9 6.056c-1.213-1.195-3.266-.336-3.266 1.366v33.156c0 1.703 2.053 2.561 3.266 1.366L40.097 24Z"/>`),
		g.Group(children),
	)
}
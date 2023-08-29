package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SsiIboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.812 30.247c.98 1.277 2.211 1.753 3.923 1.753h2.369a3.991 3.991 0 0 0 3.991-3.991v-.017A3.991 3.991 0 0 0 17.104 24H14.49a3.996 3.996 0 0 1-3.996-3.996h0A4.004 4.004 0 0 1 14.499 16h2.356c1.712 0 2.942.476 3.923 1.753m3.434 12.494c.98 1.277 2.211 1.753 3.923 1.753h2.369a3.991 3.991 0 0 0 3.991-3.991v-.017A3.991 3.991 0 0 0 30.504 24H27.89a3.996 3.996 0 0 1-3.996-3.996h0A4.004 4.004 0 0 1 27.899 16h2.356c1.712 0 2.942.476 3.923 1.753M37.505 32V16"/>`),
		g.Group(children),
	)
}
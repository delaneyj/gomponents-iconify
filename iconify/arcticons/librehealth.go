package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Librehealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.994 9.46v7.92h-7.92v6.01h7.92v7.922h6.013V23.39h7.92v-6.01h-7.92V9.46Z"/><path fill="none" stroke="currentColor" d="M40.956 32.167V6.608A2.108 2.108 0 0 0 38.849 4.5H9.151a2.108 2.108 0 0 0-2.108 2.108v25.559a3.651 3.651 0 0 0 1.826 3.162l13.305 7.682a3.651 3.651 0 0 0 3.652 0l13.305-7.682a3.651 3.651 0 0 0 1.826-3.162Z"/>`),
		g.Group(children),
	)
}
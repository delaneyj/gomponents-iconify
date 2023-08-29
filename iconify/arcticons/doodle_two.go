package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoodleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.849 16.718l3.826 7.652a3.351 3.351 0 0 1-1.149 4.215l-22.197 10.72a3.029 3.029 0 0 1-4.09-1.27q-.07-.131-.125-.269l-9.03-18.304a7.862 7.862 0 0 1 13.558-7.966l14.612 30.865m1.148-36.723a3.807 3.807 0 1 1-3.807 3.808a3.8 3.8 0 0 1 3.807-3.808Z"/>`),
		g.Group(children),
	)
}
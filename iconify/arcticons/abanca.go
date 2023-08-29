package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abanca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.076 28.789h12.937M43.5 38.74h-2.138a1.458 1.458 0 0 1-1.384-1L30.544 9.26L20.78 38.74m1.625-29.479L12.64 38.739m1.625-29.478L4.5 38.739"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caixa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 37.361l12.333-12.333H29.14L16.806 37.36H5.5Zm23.64-12.333l6.166 12.333H24l-2.35-4.701m-2.789-9.688L12.694 10.64H24l6.167 12.333"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.398 15.435l4.796-4.796H42.5L30.167 22.972H18.86"/>`),
		g.Group(children),
	)
}
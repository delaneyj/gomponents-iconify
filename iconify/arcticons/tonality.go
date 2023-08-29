package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tonality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5v26.981h-6.393V3.467m27.187 26.014h-5.149V9.253M12.981 5.535v23.946H6.589V11.385m14.215 18.096v15.782M9.785 40.13V29.481M31.822 3.968v40.064m11.019-14.551v4.883"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 23.958L32.804 38.054L20.017 21.202h0a7.076 7.076 0 1 1 11.422 0h0L18.652 38.054L5.865 21.202h0a7.08 7.08 0 0 1 7.62-10.995"/>`),
		g.Group(children),
	)
}
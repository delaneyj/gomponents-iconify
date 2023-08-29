package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullwidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 6V42M17 19L12 24M12 24L17 29M12 24H36M31 19L36 24M36 24L31 29M6 6L6 42"/>`),
		g.Group(children),
	)
}
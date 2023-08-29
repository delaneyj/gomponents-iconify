package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaiteki(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.987 38.16c0 3.024-2.452 5.475-5.478 5.475S8.03 41.183 8.03 38.16V10.116c0-3.024 2.453-5.475 5.479-5.475s5.478 2.451 5.478 5.475V38.16Zm.007-20.187l11.74-11.732c2.14-2.138 5.608-2.138 7.748 0s2.14 5.605 0 7.743l-19.48 19.467m8.191-7.733l8.579 8.574a5.472 5.472 0 0 1 0 7.742a5.48 5.48 0 0 1-7.748 0l-8.443-8.437"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleOneFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 24c0-9.665 7.835-17.5 17.5-17.5S41.5 14.335 41.5 24S33.665 41.5 24 41.5S6.5 33.665 6.5 24ZM24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Zm2.5 10.25a1.25 1.25 0 0 0-2.47-.268l-.001.001l-.001.004l-.008.034c-.008.032-.02.084-.04.153a9.687 9.687 0 0 1-.184.602a12.472 12.472 0 0 1-.861 1.968c-.83 1.518-2.137 3.103-4.117 3.833a1.25 1.25 0 1 0 .864 2.346c1.912-.704 3.312-1.954 4.318-3.25V32.75a1.25 1.25 0 1 0 2.5 0v-18.5Z"/>`),
		g.Group(children),
	)
}
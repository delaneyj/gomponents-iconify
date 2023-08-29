package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockDismissTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.25 12.5h-4a.75.75 0 0 1-.75-.75v-6a.75.75 0 0 1 1.5 0V11h3.25a.75.75 0 0 1 0 1.5ZM13 1C7.478 1 3 5.478 3 11c0 .334.016.665.049.991a6.5 6.5 0 0 1 8.96 8.96c.326.033.657.049.991.049c5.522 0 10-4.478 10-10S18.522 1 13 1Zm-1 16.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Zm-7.146-2.354a.5.5 0 0 0-.708.708L5.793 17.5l-1.647 1.646a.5.5 0 0 0 .708.708L6.5 18.207l1.646 1.647a.5.5 0 0 0 .708-.708L7.207 17.5l1.647-1.646a.5.5 0 0 0-.708-.708L6.5 16.793l-1.646-1.647Z"/>`),
		g.Group(children),
	)
}
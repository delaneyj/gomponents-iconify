package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VotingChipOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19q-2.925 0-4.963-2.038T1 12q0-2.925 2.038-4.963T8 5h8q2.925 0 4.963 2.038T23 12q0 2.925-2.038 4.963T16 19H8Zm0-2h8q2.075 0 3.538-1.463T21 12q0-2.075-1.463-3.538T16 7H8Q5.925 7 4.462 8.463T3 12q0 2.075 1.463 3.538T8 17Zm.25-2h1.5v-2.25H12v-1.5H9.75V9h-1.5v2.25H6v1.5h2.25V15Zm7.25 0H17V9h-3v1.5h1.5V15ZM12 12Z"/>`),
		g.Group(children),
	)
}
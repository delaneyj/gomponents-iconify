package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M336 192h40a40 40 0 0 1 40 40v192a40 40 0 0 1-40 40H136a40 40 0 0 1-40-40V232a40 40 0 0 1 40-40h40m160-64l-80-80l-80 80m80 193V48"/>`),
		g.Group(children),
	)
}
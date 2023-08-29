package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 4c-.523 0-1.059.184-1.438.563C7.184 4.94 7 5.476 7 6v20c0 .523.184 1.059.563 1.438c.378.378.914.562 1.437.562h14c.523 0 1.059-.184 1.438-.563c.378-.378.562-.914.562-1.437V11.406c0-.523-.277-1.09-.594-1.406L19 4.594C18.684 4.277 18.117 4 17.594 4zm0 2h8.594L23 11.406V26H9zm2 9v9h10v-9zm2 2h6v5h-6z"/>`),
		g.Group(children),
	)
}
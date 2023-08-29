package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piechart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15.583 15.917l1.648-10.78A11.294 11.294 0 0 0 15.584 5C9.553 5 4.666 9.888 4.666 15.917c0 6.03 4.888 10.917 10.917 10.917S26.5 21.946 26.5 15.917c0-.256-.02-.507-.038-.76l-10.88.76zm3.854-12.79l-1.648 10.78l10.878-.76a10.908 10.908 0 0 0-9.23-10.02z"/>`),
		g.Group(children),
	)
}
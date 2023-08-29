package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M152 0q-21 0-21 21v297l-94-77q-7-6-16-5t-14 7q-6 7-5 16t7 14l143 111l141-111q15-15 2-30q-16-14-30-2l-92 77V21q0-21-21-21z"/>`),
		g.Group(children),
	)
}
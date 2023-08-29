package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallPenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.85 11.698l-.708-.707l-9.9 9.9H3v-4.243L14.314 5.334l5.657 5.657a1 1 0 0 1 0 1.414l-7.072 7.071l-1.414-1.414l6.364-6.364Zm-2.122-2.121l-1.414-1.414L5 17.476v1.415h1.414l9.314-9.314Zm2.828-7.071l2.829 2.828a1 1 0 0 1 0 1.414L19.97 8.163L15.728 3.92l1.414-1.414a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}
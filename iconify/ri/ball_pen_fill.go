package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallPenFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.85 11.805l-.708-.707l-9.9 9.9H3v-4.243L14.314 5.44l5.657 5.657a1 1 0 0 1 0 1.414l-7.072 7.071l-1.414-1.414l6.364-6.364Zm.706-9.192l2.829 2.828a1 1 0 0 1 0 1.414L19.97 8.27l-4.243-4.243l1.414-1.414a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}
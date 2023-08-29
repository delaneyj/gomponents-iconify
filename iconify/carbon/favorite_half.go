package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FavoriteHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.21 17.061L16 29l11.79-11.939a7.731 7.731 0 0 0 0-10.823a7.494 7.494 0 0 0-10.684 0L16 7.364l-1.106-1.126a7.494 7.494 0 0 0-10.684 0a7.731 7.731 0 0 0 0 10.823Zm22.145-1.416L16 26.125V10.23q1.27-1.288 2.541-2.574a5.477 5.477 0 0 1 7.814 0a5.708 5.708 0 0 1 0 7.989Z"/>`),
		g.Group(children),
	)
}
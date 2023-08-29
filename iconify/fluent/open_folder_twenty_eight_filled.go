package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenFolderTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.75a5 5 0 0 1 5-5h12a5 5 0 0 1 5 5V13a1 1 0 1 1-2 0V7.75a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3V20a3 3 0 0 0 3 3h5a1 1 0 1 1 0 2H8a5 5 0 0 1-5-5V7.75ZM10 11a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-5.53l8.24 8.295a1 1 0 0 1-1.42 1.41L12 13.358V19a1 1 0 1 1-2 0v-8Z"/>`),
		g.Group(children),
	)
}
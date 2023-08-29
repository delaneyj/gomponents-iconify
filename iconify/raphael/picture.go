package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2.5 4.833v22.334h27V4.833h-27zM25.25 25.25H6.75V6.75h18.5v18.5zM11.25 14a2.584 2.584 0 1 0-.001-5.167A2.584 2.584 0 0 0 11.25 14zm13 2.25l-4.916-4.917l-6.917 6.917l-1.917-1.917l-2.752 2.752v5.165H24.25v-8z"/>`),
		g.Group(children),
	)
}
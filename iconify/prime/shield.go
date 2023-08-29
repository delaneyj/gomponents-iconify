package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.75a.87.87 0 0 1-.28-.05A14.27 14.27 0 0 1 3.29 6.43a.74.74 0 0 1 .61-.69a27.12 27.12 0 0 0 7.79-2.42a.75.75 0 0 1 .62 0a27.12 27.12 0 0 0 7.79 2.42a.74.74 0 0 1 .61.69a14.27 14.27 0 0 1-8.43 14.27a.87.87 0 0 1-.28.05ZM4.76 7.11A12.47 12.47 0 0 0 12 19.18a12.47 12.47 0 0 0 7.24-12.07A27.56 27.56 0 0 1 12 4.82a27.56 27.56 0 0 1-7.24 2.29Z"/>`),
		g.Group(children),
	)
}
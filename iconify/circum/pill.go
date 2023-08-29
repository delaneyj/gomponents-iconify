package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.31 4.691a5.5 5.5 0 0 0-7.78 0l-6.84 6.84a5.5 5.5 0 0 0 3.89 9.39a5.524 5.524 0 0 0 3.89-1.61l6.84-6.84a5.5 5.5 0 0 0 0-7.78Zm-.71 7.07l-3.42 3.42l-6.36-6.36L12.24 5.4a4.5 4.5 0 0 1 7.68 3.17a4.429 4.429 0 0 1-1.32 3.191Z"/>`),
		g.Group(children),
	)
}
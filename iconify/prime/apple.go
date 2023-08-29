package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.52 12.46a3.31 3.31 0 0 1 1.78-3a3.85 3.85 0 0 0-3-1.6c-1.27-.1-2.65.74-3.16.74s-1.77-.7-2.73-.7c-2 0-4.11 1.59-4.11 4.76a9 9 0 0 0 .51 2.9c.44 1.28 2.09 4.49 3.81 4.44c.9 0 1.54-.64 2.71-.64s1.72.64 2.73.64c1.73 0 3.23-2.95 3.66-4.26a3.54 3.54 0 0 1-2.2-3.28Zm-2-5.87A3.37 3.37 0 0 0 15.35 4a3.79 3.79 0 0 0-2.42 1.25A3.41 3.41 0 0 0 12 7.81a3 3 0 0 0 2.5-1.22Z"/>`),
		g.Group(children),
	)
}
package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.66 4.3a1 1 0 0 0-.98-.8H5.5a1 1 0 0 0-1 1v15a1 1 0 1 0 2 0v-6h5.6l.24 1.2c.09.468.503.805.98.8h5.18a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-4.6l-.24-1.2Z"/>`),
		g.Group(children),
	)
}
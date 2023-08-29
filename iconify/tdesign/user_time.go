package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm12 7a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM12.5 18a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM8 16a4 4 0 0 0-4 4h7.05v2H2v-2a6 6 0 0 1 6-6h3v2H8Zm11-.248v1.834L20.414 19L19 20.414l-2-2v-2.662h2Z"/>`),
		g.Group(children),
	)
}
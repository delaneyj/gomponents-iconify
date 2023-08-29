package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraduationCapBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m249.8 85.49l-116-64a12 12 0 0 0-11.6 0l-116 64a12 12 0 0 0 0 21l21.8 12v47.76a19.89 19.89 0 0 0 5.09 13.32C46.63 194.7 77 220 128 220a136.88 136.88 0 0 0 40-5.75V240a12 12 0 0 0 24 0v-35.88a119.53 119.53 0 0 0 30.91-24.51a19.89 19.89 0 0 0 5.09-13.32v-47.76l21.8-12a12 12 0 0 0 0-21ZM128 45.71L219.16 96L186 114.3a1.88 1.88 0 0 1-.18-.12l-52-28.69a12 12 0 0 0-11.6 21l39 21.49L128 146.3L36.84 96ZM128 196c-40.42 0-64.65-19.07-76-31.27v-33l70.2 38.74a12 12 0 0 0 11.6 0l34.2-18.83v37.23a110.46 110.46 0 0 1-40 7.13Zm76-31.27a93.21 93.21 0 0 1-12 10.81v-37.15l12-6.62Z"/>`),
		g.Group(children),
	)
}
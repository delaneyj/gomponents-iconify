package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointingThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 92a23.88 23.88 0 0 0-16.07 6.19A24 24 0 0 0 132 82.13V44a24 24 0 0 0-48 0v94l-11.25-18.06A24 24 0 0 0 31.2 144l4.68 8.25C53.21 182.8 64.66 203 77.66 216.33C91.28 230.3 105.86 236 128 236a84.09 84.09 0 0 0 84-84v-36a24 24 0 0 0-24-24Zm16 60a76.09 76.09 0 0 1-76 76c-40 0-51.35-20.08-85.16-79.71L38.15 140a16 16 0 0 1 27.71-16a.75.75 0 0 1 .07.12l18.68 30A4 4 0 0 0 92 152V44a16 16 0 0 1 32 0v68a4 4 0 0 0 8 0v-12a16 16 0 0 1 32 0v20a4 4 0 0 0 8 0v-4a16 16 0 0 1 32 0Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 74h-82V32a6 6 0 0 0-10.24-4.24l-96 96a6 6 0 0 0 0 8.48l96 96A6 6 0 0 0 126 224v-42h82a14 14 0 0 0 14-14V88a14 14 0 0 0-14-14Zm2 94a2 2 0 0 1-2 2h-88a6 6 0 0 0-6 6v33.51L32.49 128L114 46.49V80a6 6 0 0 0 6 6h88a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}
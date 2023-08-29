package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarricadeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 60H32a20 20 0 0 0-20 20v72a20 20 0 0 0 20 20h20v28a12 12 0 0 0 24 0v-28h104v28a12 12 0 0 0 24 0v-28h20a20 20 0 0 0 20-20V80a20 20 0 0 0-20-20Zm-4 59l-35-35h35Zm-43 29l-64-64h38l64 64Zm-72 0L41 84h38l64 64Zm-69-35l35 35H36Z"/>`),
		g.Group(children),
	)
}
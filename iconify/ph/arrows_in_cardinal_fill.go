package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInCardinalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m101.66 133.66l-32 32A8 8 0 0 1 56 160v-24H24a8 8 0 0 1 0-16h32V96a8 8 0 0 1 13.66-5.66l32 32a8 8 0 0 1 0 11.32Zm20.68-32a8 8 0 0 0 11.32 0l32-32A8 8 0 0 0 160 56h-24V24a8 8 0 0 0-16 0v32H96a8 8 0 0 0-5.66 13.66Zm11.32 52.68a8 8 0 0 0-11.32 0l-32 32A8 8 0 0 0 96 200h24v32a8 8 0 0 0 16 0v-32h24a8 8 0 0 0 5.66-13.66ZM232 120h-32V96a8 8 0 0 0-13.66-5.66l-32 32a8 8 0 0 0 0 11.32l32 32A8 8 0 0 0 200 160v-24h32a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}
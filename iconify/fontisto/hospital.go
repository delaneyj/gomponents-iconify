package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.962 0H2.036A2.037 2.037 0 0 0 0 2.036v19.928C0 23.088.912 24 2.036 24h19.929A2.035 2.035 0 0 0 24 21.965V2.037A2.036 2.036 0 0 0 21.964.002h-.001zM19.84 18.763h-4.29v-4.66H8.469v4.66h-4.32V5.219h4.32v4.66h7.081v-4.66h4.29z"/>`),
		g.Group(children),
	)
}
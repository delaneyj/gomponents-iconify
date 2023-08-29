package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeStudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.754 34.25L24 44.5L6.247 34.25v-20.5L24 3.5l17.754 10.25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.258 30.5L24 37l-11.258-6.5v-13L24 11l11.258 6.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.764 29.03l8.488-5.013l-8.504-5.047Z"/>`),
		g.Group(children),
	)
}
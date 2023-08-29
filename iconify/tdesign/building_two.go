package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .834l6.372 3.823l-1.029 1.715L17 6.166V8h5v14H2V8h5V6.166l-.343.206l-1.029-1.715L12 .834ZM9 4.966V8h6V4.966l-3-1.8l-3 1.8Zm2 .032h2.004v2.004H11V4.998ZM4 10v10h4v-7h8v7h4V10H4Zm10 10v-5h-4v5h4Z"/>`),
		g.Group(children),
	)
}
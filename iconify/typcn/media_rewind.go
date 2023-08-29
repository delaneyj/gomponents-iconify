package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaRewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.2 6.4a1.79 1.79 0 0 0-1.253.512C6.566 9.227 3 12.701 3 12.701l5.944 5.789A1.802 1.802 0 0 0 12 17.201v-9c0-.994-.806-1.801-1.8-1.801zm9 0a1.79 1.79 0 0 0-1.253.512C15.566 9.227 12 12.701 12 12.701l5.944 5.789A1.802 1.802 0 0 0 21 17.201v-9c0-.994-.806-1.801-1.8-1.801z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2V2Zm2 2v5h5V4H4Zm9-2h9v9h-9V2Zm2 2v5h5V4h-5ZM2 13h9v9H2v-9Zm2 2v5h5v-5H4Zm12.5-1.5h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3Z"/>`),
		g.Group(children),
	)
}
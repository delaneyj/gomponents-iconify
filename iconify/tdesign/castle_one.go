package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2v2h2V2h2v2h2V2h2v2h2V2h2v4h-1v8h3v-1h2v9H1v-9h2v1h3V6H5V2h2Zm1 4v14h3v-7h2v7h3V6H8Zm10 14h3v-4h-3v4ZM6 20v-4H3v4h3ZM9 7.998h2.004v2.004H9V7.998Zm4 0h2.004v2.004H13V7.998Z"/>`),
		g.Group(children),
	)
}
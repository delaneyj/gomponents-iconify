package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Town(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h5.442L11 5.913V2h5.468L22 8.638V22H2V2Zm11 2v16h7V9.362L15.532 4H13Zm-2 16V8.887L6.558 4H4v16h7ZM6 7.998h2.004v2.004H6V7.998Zm9 0h2.004v2.004H15V7.998Zm-9 4h2.004v2.004H6v-2.004Zm9 0h2.004v2.004H15v-2.004Zm-9 4h2.004v2.004H6v-2.004Zm9 0h2.004v2.004H15v-2.004Z"/>`),
		g.Group(children),
	)
}
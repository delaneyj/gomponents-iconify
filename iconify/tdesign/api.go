package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Api(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.586 12L5 7.586L6.414 9l-3 3l3 3L5 16.414L.586 12Zm7-7L12 .586L16.414 5L15 6.414l-3-3l-3 3L7.586 5ZM9 17.586l3 3l3-3L16.414 19L12 23.414L7.586 19L9 17.586Zm1.998-4.584v-2.004h2.004v2.004h-2.004ZM17.586 15l3-3l-3-3L19 7.586L23.414 12L19 16.414L17.586 15Z"/>`),
		g.Group(children),
	)
}
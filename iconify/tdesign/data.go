package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Data(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v7h1.998v-.002h2.004V11H20V4H4Zm16 9H8.002v.002H5.998V13H4v7h16v-7ZM5.998 6.5h2.004v2.004H5.998V6.5Zm0 9h2.004v2.004H5.998V15.5Z"/>`),
		g.Group(children),
	)
}
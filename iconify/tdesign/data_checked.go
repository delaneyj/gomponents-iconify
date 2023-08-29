package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v12h-2v-1H4v7h8.5v2H2V2Zm2 9h16V4H4v7Zm1.998-4.5h2.004v2.004H5.998V6.5Zm0 9h2.004v2.004H5.998V15.5Zm17.598 1.44l-5.657 5.656l-3.535-3.535l1.414-1.415l2.121 2.122l4.243-4.243l1.414 1.414Z"/>`),
		g.Group(children),
	)
}
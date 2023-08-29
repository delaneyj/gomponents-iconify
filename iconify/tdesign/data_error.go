package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v12h-2v-1H4v7h10v2H2V2Zm2 9h16V4H4v7Zm1.998-4.5h2.004v2.004H5.998V6.5Zm10.88 8.964L19 17.586l2.122-2.122l1.414 1.415L20.414 19l2.122 2.121l-1.415 1.415l-2.12-2.122l-2.122 2.121l-1.414-1.414L17.586 19l-2.121-2.121l1.414-1.415Zm-10.88.036h2.004v2.004H5.998V15.5Z"/>`),
		g.Group(children),
	)
}
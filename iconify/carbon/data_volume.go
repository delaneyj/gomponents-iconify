package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 20h-4v2h4v2h-3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h5v-8c0-1.103-.897-2-2-2zm0 8h-3v-2h3v2zm-6-6v-2h-3v-2h-2v2h-2v2h2v6c0 1.103.897 2 2 2h3v-2h-3v-6h3zm-10-2h-4v2h4v2h-3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h5v-8c0-1.103-.897-2-2-2zm0 8h-3v-2h3v2zM5 16v4H2c-1.103 0-2 .897-2 2v6c0 1.103.897 2 2 2h5V16H5zm-3 6h3v6H2v-6zm2-8V5h7.586l4 4H28v8h2V9a2.002 2.002 0 0 0-2-2H16.414L13 3.586A1.987 1.987 0 0 0 11.586 3H4a2.002 2.002 0 0 0-2 2v9h2z"/>`),
		g.Group(children),
	)
}
package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h2.93a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 12.07 6H18a3 3 0 0 1 3 3v2h-2V9a1 1 0 0 0-1-1h-5.93a3 3 0 0 1-2.496-1.336l-.812-1.219A1 1 0 0 0 7.93 5H5a1 1 0 0 0-1 1v12a1 1 0 1 0 2 0v-2h2v2a3 3 0 1 1-6 0V6z"/><path d="M6 13a3 3 0 0 1 3-3h11a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3H5v-2c.173 0 .456-.06.666-.212c.159-.114.334-.314.334-.788v-5zm1.853 6H20a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v5c0 .368-.053.701-.147 1z"/></g>`),
		g.Group(children),
	)
}
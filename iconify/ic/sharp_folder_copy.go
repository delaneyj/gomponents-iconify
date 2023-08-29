package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFolderCopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 6H1v15h19v-2H3z"/><path fill="currentColor" d="M23 4h-9l-2-2H5.01L5 17h18V4z"/>`),
		g.Group(children),
	)
}
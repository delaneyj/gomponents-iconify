package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.482 2.5l2 2.5H14v2h-3.48l-2-2.5H3v11.501l.003 1.412l-2 2L1 16.002V2.5h8.482ZM21.414 4l-1 1H23v15H5.414L3 22.414L1.586 21L20 2.586L21.414 4Zm-3 3l-11 11H21V7h-2.586Z"/>`),
		g.Group(children),
	)
}
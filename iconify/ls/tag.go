package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 2h267l379 379c44 44 0 89 0 89S512 603 468 648c-45 44-89 0-89 0L0 269V2zm183 183c22-23 22-57 0-79s-57-22-79 0s-22 57 0 79s57 22 79 0z"/>`),
		g.Group(children),
	)
}
package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M248 270v66c0 14-11 20-22 12L9 192c-12-8-12-23 0-31L226 4c11-8 22-3 22 11v65c49 0 150 6 216 49c151 99 238 423-256 587c0 0 298-149 246-354c-13-55-74-104-206-92z"/>`),
		g.Group(children),
	)
}
package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spacengine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M312.548 375.792L0 261.172v-40.987l312.548-115.95v46.596L64.47 238.052l248.078 89.958v47.782zM429.493 0v512H474V0h-44.507z"/>`),
		g.Group(children),
	)
}
package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guilsinglleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M516 556L11 324l-11-6v-67l11-7L516 14v79L98 284l418 192v80z"/>`),
		g.Group(children),
	)
}
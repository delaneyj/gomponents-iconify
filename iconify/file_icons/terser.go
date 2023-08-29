package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 177l-68-39V0h137v138l-69 39zm-68-39L69 69L0 187l120 69l68-39v-79zm204 118l120-69l-68-118l-119 69v79l67 39zm-136-79l-68 40v78l68 40l69-40v-78l-69-40zm-136 79L0 325l68 118l120-69v-79l-68-39zm205 39v79l119 69l68-118l-120-69l-67 39zm-137 79v138h137V374l-69-39l-68 39z"/>`),
		g.Group(children),
	)
}
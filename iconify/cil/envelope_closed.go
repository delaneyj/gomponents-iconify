package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 112v384h480V112Zm220.8 229.6a32.167 32.167 0 0 0 38.4 0l23.467-17.6L464 448v16H48v-16l165.333-124ZM256 316L48 160v-16h416v16ZM48 200l138.667 104L48 408Zm416 208L325.333 304L464 200Z"/>`),
		g.Group(children),
	)
}
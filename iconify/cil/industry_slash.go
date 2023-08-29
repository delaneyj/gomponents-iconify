package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndustrySlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M459.26 96L328 225.697V96h-34.55l-64.44 65.128l22.628 22.628L296 138.92v89.198l37.314 37.315L464 136.303v259.815l32 32V96h-36.74zM168 16H83.882L168 100.118V16zm-32 288h32v32h-32zm0 72h32v32h-32zm80-72h32v32h-32zm0 72h32v32h-32zm80 0h32v32h-32z"/><path fill="currentColor" d="M38.627 16H16v480h480v-22.627ZM48 464V70.627l88 88V248h32v-57.373L441.373 464Z"/>`),
		g.Group(children),
	)
}
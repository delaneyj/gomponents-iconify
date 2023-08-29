package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiggAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 93h-59V3H231v89h-59v116H32V92H2v117h30v31h54v54H63v32h77v55H63v-55H32v87h139v-87h-31v-32h-23v-54h87v172h232V233h26V93zm-91-59v114H262V34h109zm-36 347v-87h-28v87h-71V208h-33v-83h28v55h172v-55h27v84h-25v172h-70zM140 92V62h32v30h-32zm0-30H63V30h77v32zM63 92H32V62h31v30z"/>`),
		g.Group(children),
	)
}
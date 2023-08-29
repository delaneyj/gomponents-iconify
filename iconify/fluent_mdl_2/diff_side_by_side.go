package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiffSideBySide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h896v1280H0V384zm768 1152V768H128v768h640zM128 512v128h640V512H128zm1920-128v1280h-896V384h896zm-128 896h-640v256h640v-256zm0-768h-640v512h640V512z"/>`),
		g.Group(children),
	)
}
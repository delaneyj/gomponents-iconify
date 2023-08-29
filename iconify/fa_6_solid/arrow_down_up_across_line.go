package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownUpAcrossLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M137.4 502.6c12.5 12.5 32.8 12.5 45.3 0l96-96c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L192 402.7V288h352c17.7 0 32-14.3 32-32s-14.3-32-32-32h-96V109.3l41.4 41.4c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3l-96-96c-12.5-12.5-32.8-12.5-45.3 0l-96 96c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0l41.3-41.4V224H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h96v114.7l-41.4-41.3c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l96 96zM128 192h64V64c0-17.7-14.3-32-32-32s-32 14.3-32 32v128zm320 128h-64v128c0 17.7 14.3 32 32 32s32-14.3 32-32V320z"/>`),
		g.Group(children),
	)
}
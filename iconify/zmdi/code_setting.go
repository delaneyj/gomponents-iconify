package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M148 80L55 192l93 112l-33 27L0 192L115 53zm-16 133v-42h43v42h-43zm213-42v42h-42v-42h42zm-128 42v-42h43v42h-43zM362 53l115 139l-115 139l-33-27l93-112l-93-112z"/>`),
		g.Group(children),
	)
}
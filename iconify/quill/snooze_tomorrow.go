package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnoozeTomorrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21h22M7 26h18M10 16a6 6 0 0 1 12 0M6.5 16h-2M9 9L7.5 7.5M16 6V4m7 5l1.485-1.485M28 16h-2.5"/>`),
		g.Group(children),
	)
}
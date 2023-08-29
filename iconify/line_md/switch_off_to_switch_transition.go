package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchOffToSwitchTransition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"/><circle cx="7" cy="12" r="3" fill="currentColor"><animate fill="freeze" attributeName="cx" dur="0.2s" values="7;17"/></circle>`),
		g.Group(children),
	)
}
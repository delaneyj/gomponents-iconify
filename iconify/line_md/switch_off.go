package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><circle cx="7" cy="12" r="3" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></circle>`),
		g.Group(children),
	)
}
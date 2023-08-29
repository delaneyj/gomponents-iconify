package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchOffFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdSwitchOffFilled0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="8" x="3" y="8" stroke-dasharray="54" stroke-dashoffset="54" rx="4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></g><circle cx="7" cy="12" r="3" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></mask><rect width="20" height="10" x="2" y="7" fill="currentColor" mask="url(#lineMdSwitchOffFilled0)"/>`),
		g.Group(children),
	)
}
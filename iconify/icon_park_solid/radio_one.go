package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRadioOne0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="38" height="28" x="5" y="14" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 22h4m-4 6h4m-4 6h4"/><circle cx="18" cy="28" r="7" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 14V6h28v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRadioOne0)"/>`),
		g.Group(children),
	)
}
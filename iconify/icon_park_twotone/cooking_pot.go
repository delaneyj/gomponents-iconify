package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookingPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCookingPot0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M10 44h28V20.947C38 14.901 31.732 10 24 10s-14 4.901-14 10.947V44Z" clip-rule="evenodd"/><path fill="#555" d="M38 22.044v-1.097C38 14.901 31.732 10 24 10s-14 4.901-14 10.947v1.093l28 .004Z"/><path stroke-linecap="round" d="M4 22h40M21 4h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCookingPot0)"/>`),
		g.Group(children),
	)
}
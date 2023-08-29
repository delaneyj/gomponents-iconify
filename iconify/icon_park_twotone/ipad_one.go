package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IpadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIpadOne0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="38" height="30" x="5" y="10" fill="#555" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M11 27v-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIpadOne0)"/>`),
		g.Group(children),
	)
}
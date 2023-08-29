package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPlugOne0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="24" height="24" x="12" y="12" fill="#555" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 36v5a3 3 0 0 1-3 3H8m12-32V4m8 8V4m-6 20h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPlugOne0)"/>`),
		g.Group(children),
	)
}
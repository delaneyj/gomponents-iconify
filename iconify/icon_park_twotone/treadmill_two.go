package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreadmillTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTreadmillTwo0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M8 39v5m32-5v5m-4-13l6-16l-4-6m-5 5L43 4"/><rect width="40" height="8" x="4" y="31" fill="#555" rx="4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTreadmillTwo0)"/>`),
		g.Group(children),
	)
}
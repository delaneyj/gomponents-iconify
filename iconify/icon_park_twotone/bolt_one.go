package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBoltOne0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M12 12.28a.28.28 0 0 1 .28-.28h23.44a.28.28 0 0 1 .28.28V24c0 6.627-5.373 12-12 12s-12-5.373-12-12V12.28Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 12V4m8 8V4m-6 23h4m-2 9v5a3 3 0 0 0 3 3h11"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBoltOne0)"/>`),
		g.Group(children),
	)
}
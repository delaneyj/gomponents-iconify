package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPicture0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path fill="#555" d="M18 23a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path d="M27.79 26.22a2 2 0 0 1 3.243.053l8.775 12.583c.924 1.326-.025 3.144-1.64 3.144H16l11.79-15.78Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPicture0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOven0"><g fill="none"><rect width="40" height="28" x="4" y="8" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><rect width="16" height="12" x="12" y="16" fill="#000" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="1"/><circle cx="37" cy="15" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 22h2m-2 7h2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 36v6m24-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOven0)"/>`),
		g.Group(children),
	)
}
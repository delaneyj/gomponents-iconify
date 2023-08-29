package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videocamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVideocamera0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="32" height="26" x="4" y="11" rx="2"/><circle cx="20" cy="24" r="6" fill="#fff" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="m36 29l8 4V15l-8 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVideocamera0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSArchery0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M13 42c9.941 0 18-8.059 18-18S22.941 6 13 6"/><circle cx="9" cy="24" r="3" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 24h30m-4-4l4 4l-4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSArchery0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EveryUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEveryUser0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="14" cy="29" r="5" fill="#fff"/><circle cx="34" cy="29" r="5" fill="#fff"/><circle cx="24" cy="9" r="5" fill="#fff"/><path d="M24 44c0-5.523-4.477-10-10-10S4 38.477 4 44m40 0c0-5.523-4.477-10-10-10s-10 4.477-10 10m10-20c0-5.523-4.477-10-10-10s-10 4.477-10 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEveryUser0)"/>`),
		g.Group(children),
	)
}
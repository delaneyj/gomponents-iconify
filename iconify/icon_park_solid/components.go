package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Components(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSComponents0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17 11l7-7l7 7l-7 7l-7-7Zm13 14l7-7l7 7l-7 7l-7-7ZM17 37l7-7l7 7l-7 7l-7-7ZM4 24l7-7l7 7l-7 7l-7-7Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSComponents0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrongUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWrongUser0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="12" r="8" fill="#fff"/><path d="M42 44c0-9.941-8.059-18-18-18S6 34.059 6 44m14-8l8 8m0-8l-8 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWrongUser0)"/>`),
		g.Group(children),
	)
}
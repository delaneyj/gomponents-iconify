package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intercom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIntercom0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M13 14a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v11l-2 6v10a3 3 0 0 1-3 3H18a3 3 0 0 1-3-3V31l-2-6V14Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M19 11V4m9 7V7"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 19h-8m6 7h-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIntercom0)"/>`),
		g.Group(children),
	)
}
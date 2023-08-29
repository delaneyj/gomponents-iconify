package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Avocado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAvocado0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M33 19c-2-5 0-15-9-15s-7 10-9 14s-5 7-5 13c0 9 7.5 13 14 13s14-4 14-13c0-5-3-7-5-12Z"/><path fill="#000" stroke="#000" d="M30 30a6 6 0 0 1-12 0c0-3.314 2.686-7.5 6-7.5s6 4.186 6 7.5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAvocado0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHotel0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M4 4h40"/><rect width="32" height="40" x="8" y="4" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M20 32h8v12h-8V32Z"/><path stroke="#000" stroke-linecap="round" d="M15 12h2m-2 6h2m6-6h2m-2 6h2m6-6h2m-2 6h2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M4 44h40"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 32h2c.552 0 1.01-.452.904-.994C30.352 28.166 27.471 26 24 26c-3.47 0-6.352 2.165-6.904 5.006c-.106.542.352.994.904.994h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHotel0)"/>`),
		g.Group(children),
	)
}
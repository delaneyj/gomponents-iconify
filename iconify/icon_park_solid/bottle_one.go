package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBottleOne0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M15 30a9 9 0 0 1 1.8-5.4l3.6-4.8A3 3 0 0 0 21 18V4h6v14a3 3 0 0 0 .6 1.8l3.6 4.8A9 9 0 0 1 33 30v12a2 2 0 0 1-2 2H17a2 2 0 0 1-2-2V30Z"/><path stroke="#000" d="M21 10h6"/><path stroke="#fff" d="M21 12V8m6 4V8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBottleOne0)"/>`),
		g.Group(children),
	)
}
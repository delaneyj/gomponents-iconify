package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCloudy0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M9.455 30.994A13.95 13.95 0 0 1 4 19.885C4 12.217 10.105 6 17.636 6c6.297 0 11.598 4.346 13.166 10.253a8.921 8.921 0 0 1 4.107-.996c5.02 0 9.091 4.144 9.091 9.257c0 3.795-2.244 7.058-5.455 8.486"/><path fill="#fff" d="M26 42a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path stroke-linecap="round" d="M22.243 15.757a6 6 0 0 0-8.485 8.485"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCloudy0)"/>`),
		g.Group(children),
	)
}
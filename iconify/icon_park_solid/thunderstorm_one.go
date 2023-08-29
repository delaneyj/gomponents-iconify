package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThunderstormOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSThunderstormOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M20.53 20L17 31.5l6.177.885L20.529 43L32 29.73l-7.059-1.768L30.235 20H20.53Z"/><path stroke-linecap="round" d="M9.455 29.994A13.95 13.95 0 0 1 4 18.885C4 11.217 10.105 5 17.636 5c6.297 0 11.598 4.346 13.166 10.253a8.921 8.921 0 0 1 4.107-.996c5.02 0 9.091 4.144 9.091 9.257c0 3.795-2.244 7.058-5.455 8.486"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSThunderstormOne0)"/>`),
		g.Group(children),
	)
}
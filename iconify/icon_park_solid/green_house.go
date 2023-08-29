package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGreenHouse0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M42 20v24H6V20L24 4l18 16Z"/><path stroke-linecap="round" d="M6 24h36M13 14v30m22-30v30"/><path fill="#fff" stroke-linecap="round" d="M20 32h8v12h-8z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGreenHouse0)"/>`),
		g.Group(children),
	)
}
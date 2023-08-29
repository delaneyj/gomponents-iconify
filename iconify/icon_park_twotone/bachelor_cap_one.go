package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BachelorCapOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBachelorCapOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" stroke-linecap="round" d="M5 16L24 6l19 10l-19 10L5 16Z"/><path d="M11 20v14.464c0 1.464 1.055 2.723 2.471 3.095c2.205.58 5.585 1.66 8.885 3.47c1.021.56 2.266.56 3.288 0c3.3-1.81 6.68-2.89 8.885-3.47C35.945 37.187 37 35.93 37 34.464V20"/><path stroke-linecap="round" d="M43 16v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBachelorCapOne0)"/>`),
		g.Group(children),
	)
}
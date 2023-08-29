package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinMiniApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWeixinMiniApp0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path d="M29 24a5 5 0 1 0-5-5v10a5 5 0 1 1-5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWeixinMiniApp0)"/>`),
		g.Group(children),
	)
}
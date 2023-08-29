package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BytedanceMiniApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBytedanceMiniApp0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" d="m26.122 21.879l-4.243 4.242m10.606.707l1.414-1.414a4 4 0 0 0 0-5.657L28.242 14.1a4 4 0 0 0-5.656 0l-1.415 1.415m5.657 16.97L25.414 33.9a4 4 0 0 1-5.657 0L14.1 28.243a4 4 0 0 1 0-5.657l1.415-1.414"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBytedanceMiniApp0)"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200.73 36H55.27A19.3 19.3 0 0 0 36 55.27v145.46A19.3 19.3 0 0 0 55.27 220h145.46A19.3 19.3 0 0 0 220 200.73V55.27A19.3 19.3 0 0 0 200.73 36ZM196 196H60V60h136Z"/>`),
		g.Group(children),
	)
}
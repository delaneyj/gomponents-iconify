package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16v9h2v11H2V11h2V2Zm0 11v7h16v-7H4Zm14-2V4H6v7h12ZM8 6.496h2.004V8.5H8V6.496Zm6 0h2.004V8.5H14V6.496Z"/>`),
		g.Group(children),
	)
}
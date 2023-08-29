package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActivateSubscriptionsOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.004 5.999h16v2h-16zm2-3.999h12v2h-12zm13.993 7.991H5.003A2.008 2.008 0 0 0 3 12.004v7.974a2.008 2.008 0 0 0 2.003 2.013h15.994A2.008 2.008 0 0 0 23 19.978v-7.974a2.008 2.008 0 0 0-2.003-2.013ZM21 20H5v-8h16Z"/><path fill="currentColor" d="M12.004 18.991h2v-1.989h2v-2.004h-2v-2.007h-2v2.01h-2v2.001h2v1.989z"/>`),
		g.Group(children),
	)
}
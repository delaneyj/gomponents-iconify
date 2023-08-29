package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudComputingOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6.147a4.973 4.973 0 0 0-.65-.107a7.492 7.492 0 0 0-14-2A5.904 5.904 0 0 0 4 4.365A5.98 5.98 0 0 0 4 15.65v-2.204a3.976 3.976 0 0 1 0-6.901a3.93 3.93 0 0 1 1.56-.515l1.07-.11l.5-.95a5.487 5.487 0 0 1 10.26 1.46l.3 1.5l1.53.11a2.913 2.913 0 0 1 .78.171a2.963 2.963 0 0 1 0 5.604v2.084a4.972 4.972 0 0 0 0-9.752Z"/><path fill="currentColor" d="M6 10v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5Zm2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5ZM6 17v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5Zm2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5Z"/>`),
		g.Group(children),
	)
}
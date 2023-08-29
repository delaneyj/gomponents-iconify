package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudComputing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 10v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5Zm2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5ZM6 17v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5Zm2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5Z"/><path fill="currentColor" d="M22.965 8a5.35 5.35 0 0 0-3.615-1.96a7.492 7.492 0 0 0-14-2A5.904 5.904 0 0 0 4 4.365A5.98 5.98 0 0 0 4 15.65V8h16v7.9a5.003 5.003 0 0 0 4-4.9a4.908 4.908 0 0 0-1.035-3Z"/>`),
		g.Group(children),
	)
}
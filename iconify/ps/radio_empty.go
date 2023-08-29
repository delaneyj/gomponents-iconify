package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M301 192q0-62-43.5-105.5T152 43T46.5 86.5T3 192t43.5 105.5T152 341t105.5-43.5T301 192zm-256 0q0-45 31-76t76-31t76 31t31 76t-31 76t-76 31t-76-31t-31-76z"/>`),
		g.Group(children),
	)
}
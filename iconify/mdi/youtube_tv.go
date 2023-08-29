package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 4.5h19c.84 0 1.5.65 1.5 1.5v11.5c0 .85-.66 1.5-1.5 1.5h-19c-.85 0-1.5-.65-1.5-1.5V6c0-.85.65-1.5 1.5-1.5m7.21 4V15l5.71-3.3l-5.71-3.2M17.25 21H6.65c-.3 0-.5-.2-.5-.5s.2-.5.5-.5h10.7c.3 0 .5.2.5.5s-.3.5-.6.5Z"/>`),
		g.Group(children),
	)
}
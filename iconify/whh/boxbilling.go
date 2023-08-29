package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boxbilling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m960 768l-448 256L64 768L0 192L512 0l512 192zM512 192L256 320l32 256l224 128l224-128l32-256z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAddLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="8.5" cy="10.5" r="3.5" fill="currentColor"/><path fill="currentColor" d="M14 0v4h-4v2h4v4h2V6h4V4h-4V0h-2zM8 15c-4.6 0-7 2.69-7 4.23V20h14v-.77C15 17.69 12.6 15 8 15z"/>`),
		g.Group(children),
	)
}
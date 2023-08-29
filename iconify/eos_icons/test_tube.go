package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="13" cy="11" r="1" fill="currentColor"/><circle cx="11.5" cy="7.5" r=".5" fill="currentColor"/><circle cx="12.5" cy="5.5" r=".5" fill="currentColor"/><path fill="currentColor" d="M5 0v2h2v17a5 5 0 0 0 10 0V2h2V0Zm10 2v13h-3a1 1 0 0 0-2 0H9V2Z"/>`),
		g.Group(children),
	)
}
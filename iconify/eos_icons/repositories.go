package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repositories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 23.01l2.5-1.5l2.5 1.5V19H8v4.01zM8 5h2v2H8zm0 5h2v2H8zm0 5.01h2v2H8z"/><path fill="currentColor" d="M18.36 1.05A1.47 1.47 0 0 0 18 1H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2V3h12v16h-3v2h3a2 2 0 0 0 2-2V3a2 2 0 0 0-1.64-1.95Z"/>`),
		g.Group(children),
	)
}
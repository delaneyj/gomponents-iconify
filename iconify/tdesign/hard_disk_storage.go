package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDiskStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4h-2.5v7.875h-11V4H4Zm4.5 0v5.875h7V4h-7ZM14 6v3h-2V6h2Z"/>`),
		g.Group(children),
	)
}
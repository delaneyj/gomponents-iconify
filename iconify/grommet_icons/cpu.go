package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cpu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 18h3m-3-4h3m-3-4h3M1 6h3m16 12h3m-3-4h3m-3-4h3m-3-4h3M6 1v3m4-3v3m4-3v3m4-3v3M6 20v3m4-3v3m4-3v3m4-3v3M5 20h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1zm8-13h4v4h-4V7z"/>`),
		g.Group(children),
	)
}
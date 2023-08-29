package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveCage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 23V2h22v21M1 8h22H1Zm0 6h22H1Zm0 6h22H1ZM4 5h12H4Zm14 0h2h-2Zm0 6h2h-2Zm0 6h2h-2ZM4 11h12H4Zm0 6h12H4Z"/>`),
		g.Group(children),
	)
}
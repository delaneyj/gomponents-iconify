package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cli(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h22v22H1V1Zm0 4h22M5 1v4m6 11h8M5 10l3 3l-3 3"/>`),
		g.Group(children),
	)
}
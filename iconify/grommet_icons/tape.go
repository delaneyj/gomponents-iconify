package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2 2h20v20H2V2Zm17 10a7 7 0 1 1-14 0a7 7 0 0 1 14 0Zm-7-3c-1.655 0-3 1.345-3 3s1.345 3 3 3s3-1.345 3-3s-1.345-3-3-3Z"/>`),
		g.Group(children),
	)
}
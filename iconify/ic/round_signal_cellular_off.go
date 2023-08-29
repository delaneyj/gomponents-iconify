package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundSignalCellularOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3.41c0-.89-1.08-1.34-1.71-.71l-6.6 6.6L21 17.61V3.41zm.44 17.47L5.62 5.06a.996.996 0 1 0-1.41 1.41l5.66 5.66l-7.16 7.16c-.63.63-.19 1.71.7 1.71h15.32l1.29 1.29c.39.39 1.02.39 1.41 0c.4-.39.4-1.02.01-1.41z"/>`),
		g.Group(children),
	)
}
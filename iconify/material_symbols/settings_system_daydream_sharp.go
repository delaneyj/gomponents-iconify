package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsSystemDaydreamSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16h6.5q1.05 0 1.775-.725T18 13.5q0-1.05-.725-1.775T15.5 11h-.05q-.2-1.275-1.113-2.138T12.15 8q-1.05 0-1.95.525T8.85 10h-.125q-1.175 0-1.95.9T6 13q0 1.25.875 2.125T9 16Zm-7 4V4h20v16H2Z"/>`),
		g.Group(children),
	)
}
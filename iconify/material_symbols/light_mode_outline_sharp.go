package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightModeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm0 2q-2.075 0-3.538-1.463T7 12q0-2.075 1.463-3.538T12 7q2.075 0 3.538 1.463T17 12q0 2.075-1.463 3.538T12 17ZM1 13v-2h4v2H1Zm18 0v-2h4v2h-4Zm-8-8V1h2v4h-2Zm0 18v-4h2v4h-2ZM6.35 7.75L3.875 5.275l1.4-1.4L7.75 6.35l-1.4 1.4Zm12.375 12.375L16.25 17.65l1.4-1.4l2.475 2.475l-1.4 1.4ZM17.65 7.75l-1.4-1.4l2.475-2.475l1.4 1.4L17.65 7.75ZM5.275 20.125l-1.4-1.4L6.35 16.25l1.4 1.4l-2.475 2.475ZM12 12Z"/>`),
		g.Group(children),
	)
}
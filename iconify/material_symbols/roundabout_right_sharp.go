package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundaboutRightSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-6.15q-2.125-.35-3.563-1.987T2 9q0-2.5 1.75-4.25T8 3q2.225 0 3.863 1.438T13.85 8h4.325L16.6 6.4L18 5l4 4l-4 4l-1.425-1.4l1.6-1.6H11.95V9q0-1.65-1.15-2.825T8 5Q6.35 5 5.175 6.175T4 9q0 1.65 1.175 2.8T8 12.95h1V21H7Z"/>`),
		g.Group(children),
	)
}
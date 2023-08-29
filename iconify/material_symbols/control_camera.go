package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q-1.25 0-2.125-.875T9 12q0-1.25.875-2.125T12 9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15Zm0 7l-4.25-4.25l1.4-1.4l2.85 2.8l2.85-2.8l1.4 1.4L12 22Zm-5.75-5.75L2 12l4.25-4.25l1.4 1.4L4.85 12l2.8 2.85l-1.4 1.4Zm2.9-8.6l-1.4-1.4L12 2l4.25 4.25l-1.4 1.4L12 4.85l-2.85 2.8Zm8.6 8.6l-1.4-1.4l2.8-2.85l-2.8-2.85l1.4-1.4L22 12l-4.25 4.25Z"/>`),
		g.Group(children),
	)
}
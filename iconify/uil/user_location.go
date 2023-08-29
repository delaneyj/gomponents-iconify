package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.49 10.86a3.09 3.09 0 1 0-5 0a4.67 4.67 0 0 0-1.12 1A1 1 0 1 0 10 13.12a2.62 2.62 0 0 1 2.05-1a2.62 2.62 0 0 1 2.05 1a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.62a4.67 4.67 0 0 0-1.17-1.01ZM12 10.13A1.09 1.09 0 1 1 13.09 9A1.09 1.09 0 0 1 12 10.13Zm8.46-.5A8.5 8.5 0 0 0 7.3 3.36a8.56 8.56 0 0 0-3.76 6.27A8.46 8.46 0 0 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83Zm-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24Z"/>`),
		g.Group(children),
	)
}
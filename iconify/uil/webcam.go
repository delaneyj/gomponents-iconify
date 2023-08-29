package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webcam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13a3 3 0 1 0-3-3a3 3 0 0 0 3 3Zm0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm9.59 9.16A10 10 0 0 0 19 13.89a8 8 0 1 0-14 0a9.9 9.9 0 0 0-2.6 4.27a3 3 0 0 0 .47 2.63A3 3 0 0 0 5.3 22h13.4a3 3 0 0 0 2.42-1.21a3 3 0 0 0 .47-2.63ZM12 4a6 6 0 1 1-6 6a6 6 0 0 1 6-6Zm7.52 15.59a1 1 0 0 1-.82.41H5.3a1 1 0 0 1-.82-.41a1 1 0 0 1-.15-.87a7.85 7.85 0 0 1 1.88-3.22a8 8 0 0 0 11.58 0a7.85 7.85 0 0 1 1.88 3.22a1 1 0 0 1-.15.87Z"/>`),
		g.Group(children),
	)
}
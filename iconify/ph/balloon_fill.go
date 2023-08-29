package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalloonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 16a88.1 88.1 0 0 0-88 88c0 23.43 9.4 49.42 25.13 69.5c12.08 15.41 26.5 26 41.91 31.09l-10.39 24.26A8 8 0 0 0 104 240h48a8 8 0 0 0 7.35-11.15L149 204.59c15.4-5.07 29.83-15.68 41.91-31.09C206.6 153.42 216 127.43 216 104a88.1 88.1 0 0 0-88-88Zm49.32 87.89a8.52 8.52 0 0 1-1.32.11a8 8 0 0 1-7.88-6.68a41.29 41.29 0 0 0-33.43-33.43a8 8 0 1 1 2.64-15.78a57.5 57.5 0 0 1 46.57 46.57a8 8 0 0 1-6.58 9.21Z"/>`),
		g.Group(children),
	)
}
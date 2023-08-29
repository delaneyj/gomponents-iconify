package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FanBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236.85 134a64 64 0 0 0-87.43-42.11l14.22-56.77a12 12 0 0 0-5.17-13A64 64 0 0 0 86 127.52l-56.28 16.07a12 12 0 0 0-8.69 11a64 64 0 0 0 127.56 10l42 40.7a12 12 0 0 0 13.86 2a64 64 0 0 0 32.4-73.29ZM116 128a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm8-92a40 40 0 0 1 14 2.53l-13.4 53.64a36 36 0 0 0-25.85 14.88A40 40 0 0 1 124 36Zm-19 156.1A40 40 0 0 1 45.5 164l53.18-15.19a36 36 0 0 0 25.8 15A39.84 39.84 0 0 1 105 192.1Zm104.7-21.56a39.92 39.92 0 0 1-9.21 10.89L160.73 143a35.9 35.9 0 0 0 .05-29.83a40 40 0 0 1 48.89 57.4Z"/>`),
		g.Group(children),
	)
}
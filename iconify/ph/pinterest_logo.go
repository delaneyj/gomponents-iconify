package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinterestLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 112c0 22.57-7.9 43.2-22.23 58.11C180.39 184 162.25 192 144 192c-17.88 0-29.82-5.86-37.43-12l-10.78 45.83A8 8 0 0 1 88 232a8.24 8.24 0 0 1-1.84-.21a8 8 0 0 1-6-9.62l32-136a8 8 0 0 1 15.58 3.66l-16.9 71.8C114 166 123.3 176 144 176c27.53 0 56-23.94 56-64a72 72 0 1 0-134.37 36a8 8 0 0 1-13.85 8A88 88 0 1 1 216 112Z"/>`),
		g.Group(children),
	)
}
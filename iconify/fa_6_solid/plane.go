package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M482.3 192c34.2 0 93.7 29 93.7 64c0 36-59.5 64-93.7 64H365.7L265.2 495.9c-5.7 10-16.3 16.1-27.8 16.1h-56.2c-10.6 0-18.3-10.2-15.4-20.4l49-171.6H112l-43.2 57.6c-3 4-7.8 6.4-12.8 6.4H14c-7.8 0-14-6.3-14-14c0-1.3.2-2.6.5-3.9L32 256L.5 145.9c-.4-1.3-.5-2.6-.5-3.9c0-7.8 6.3-14 14-14h42c5 0 9.8 2.4 12.8 6.4L112 192h102.9l-49-171.6c-3-10.2 4.7-20.4 15.3-20.4h56.2c11.5 0 22.1 6.2 27.8 16.1L365.7 192h116.6z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShakerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.88 4l2.15 2.1l-5.53 4.4l-1-1L16.87 4h.01m0-2a2 2 0 0 0-1.55.72L9.8 9.65l3.54 3.54l6.94-5.52c.9-.76.97-2.13.13-2.97L18.3 2.59c-.4-.4-.91-.59-1.42-.59M9.1 10.36l-.71.71a1.02 1.02 0 0 0 0 1.43l2.11 2.1c.21.2.46.29.72.29s.51-.09.71-.29l.7-.7l-3.53-3.54M6 15c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1m3 1c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1m-5 2c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1m3 1c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1Z"/>`),
		g.Group(children),
	)
}
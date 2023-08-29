package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dvd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#c28fef"/><circle cx="32" cy="32" r="6.6" fill="#fff"/><path fill="#ffc7ce" d="m32.2 42.6l9.2 13.6c6.8-2.6 12.2-8 14.8-14.8l-13.6-9.2c-.1 5.7-4.7 10.3-10.4 10.4m-.4-21.2L22.6 7.8c-6.8 2.6-12.2 8-14.8 14.8l13.6 9.2c.1-5.7 4.7-10.3 10.4-10.4"/><path fill="#9450e0" d="M32 45c-7.1 0-13-5.8-13-13c0-7.1 5.8-13 13-13s13 5.8 13 13c0 7.1-5.9 13-13 13m0-23.3c-5.7 0-10.3 4.6-10.3 10.3S26.3 42.3 32 42.3c5.7 0 10.3-4.6 10.3-10.3S37.7 21.7 32 21.7"/>`),
		g.Group(children),
	)
}
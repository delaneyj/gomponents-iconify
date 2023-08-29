package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18.54 8.211a1.75 1.75 0 0 1 .577 2.407l-1.892 3.085a5.252 5.252 0 0 1-1.069 10.39a5.25 5.25 0 0 1-4.731-7.527c.042-.133.1-.263.177-.387l4.531-7.39a1.75 1.75 0 0 1 2.407-.578Zm-2.384 8.883a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}
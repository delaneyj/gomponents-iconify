package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cmake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#064F8C" d="M62.8.4L.3 123.8l68.1-57.9z"/><path fill="#249847" d="m123.8 127.7l-84-33.9L0 127.7z"/><path fill="#BE2128" d="M128 126.6L65.6 2.5l9.2 102.6z"/><path fill="#CDCDCE" d="m71.9 104l-3.1-34.9L42 92z"/>`),
		g.Group(children),
	)
}
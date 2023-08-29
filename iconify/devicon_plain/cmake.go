package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cmake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M62.8.4L.3 123.8l68.1-57.9zm61 127.3l-84-33.9L0 127.7zm4.2-1.1L65.6 2.5l9.2 102.6zM71.9 104l-3.1-34.9L42 92z"/>`),
		g.Group(children),
	)
}
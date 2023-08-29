package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poetry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M16.215 5.332v95.61h.002L40.07 124.8a121.72 121.72 0 0 0 62.242-33.367L83.819 72.939L16.252 5.367l69.35 65.738c16.25-17.138 26.222-40.29 26.222-65.773H16.215z"/>`),
		g.Group(children),
	)
}
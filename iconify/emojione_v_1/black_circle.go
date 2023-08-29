package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#354a54" d="M57.55 14.451c9.694 14.11 6.109 33.406-7.999 43.1c-14.11 9.695-33.408 6.112-43.1-7.999c-9.693-14.11-6.112-33.41 8-43.1c14.11-9.692 33.408-6.109 43.1 8"/>`),
		g.Group(children),
	)
}
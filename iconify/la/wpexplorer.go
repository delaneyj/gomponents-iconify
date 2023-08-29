package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wpexplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.8 3 3 8.8 3 16s5.8 13 13 13s13-5.8 13-13S23.2 3 16 3zm0 2c6.1 0 11 4.9 11 11s-4.9 11-11 11S5 22.1 5 16S9.9 5 16 5zm-4.8 5l-1.8 4.3l4.3 1.8l1.8-4.3l-4.3-1.8zm4.8 2.6L14.6 16l3.3 1.4l.016-.035l.084.035l1.5-3.4l-3.5-1.4zm4 2.1l-1.1 2.7l2.7 1.1l1.2-2.7l-2.8-1.1zm-5.7 1.9l-.3.6l1.1.5l-2.5 5.3h.7l2.4-4.9l2.5 4.9h.7l-2.3-4.6l.8.3l.3-.7l-3.4-1.4z"/>`),
		g.Group(children),
	)
}
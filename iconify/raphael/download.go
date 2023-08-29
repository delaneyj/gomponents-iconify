package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534zm0 27.326a2.806 2.806 0 1 1 0-5.611a2.806 2.806 0 0 1 0 5.611zm0-7.705l-7.858-6.562h3.47V5.747h8.778v8.778h3.468L16 21.087z"/>`),
		g.Group(children),
	)
}
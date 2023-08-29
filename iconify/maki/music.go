package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.5 0a.49.49 0 0 0-.23.06L3.5 2.5A.5.5 0 0 0 3 3v6.28A2 2 0 0 0 2 9a2 2 0 1 0 2 2V6.36l8-2.22v3.64a2 2 0 0 0-1-.28a2 2 0 1 0 2 2v-9a.5.5 0 0 0-.5-.5zM12 3.14L4 5.36v-2l8-2.22z"/>`),
		g.Group(children),
	)
}
package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 5l-.313.281L4.281 16.813l-.687.687l.687.719l9.5 9.5l.719.687l.688-.687l11.53-11.407L27 16V5zm.844 2H25v8.156L14.5 25.594L6.406 17.5zM22 9c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1z"/>`),
		g.Group(children),
	)
}
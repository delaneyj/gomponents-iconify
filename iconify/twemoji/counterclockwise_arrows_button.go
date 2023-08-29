package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CounterclockwiseArrowsButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="m22.242 22.242l2.829 2.829c-3.905 3.905-10.237 3.904-14.143-.001a9.968 9.968 0 0 1-2.854-8.225l-4.037.367c-.215 3.84 1.128 7.752 4.062 10.687c5.467 5.467 14.333 5.468 19.799 0l2.828 2.828l.849-9.334l-9.333.849zM27.899 8.1C22.431 2.633 13.568 2.633 8.1 8.1L5.272 5.272l-.849 9.334l9.334-.849l-2.829-2.829c3.906-3.905 10.236-3.905 14.142 0a9.975 9.975 0 0 1 2.856 8.226l4.036-.366c.216-3.841-1.128-7.753-4.063-10.688z"/>`),
		g.Group(children),
	)
}
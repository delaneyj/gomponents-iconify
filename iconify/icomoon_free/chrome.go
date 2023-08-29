package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chrome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m4.036 6.977l-2.29-3.966A7.986 7.986 0 0 1 8-.001a7.994 7.994 0 0 1 6.883 3.922H8.355a4.1 4.1 0 0 0-4.32 3.055zm6.828-1.899h4.585a8 8 0 0 1-7.358 10.921l3.272-5.667a4.08 4.08 0 0 0-.499-5.254zM5.094 8c0-1.603 1.304-2.906 2.906-2.906S10.906 6.398 10.906 8S9.602 10.906 8 10.906S5.094 9.602 5.094 8zm4.003 3.944l-2.29 3.967a8.001 8.001 0 0 1-5.78-11.833l3.266 5.657a4.1 4.1 0 0 0 4.804 2.21z"/>`),
		g.Group(children),
	)
}
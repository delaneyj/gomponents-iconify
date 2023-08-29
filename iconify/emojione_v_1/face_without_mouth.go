package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithoutMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fbbf67" d="M63.89 31.945c0 17.641-14.301 31.943-31.945 31.943C14.299 63.889 0 49.586 0 31.945C0 14.301 14.299 0 31.945 0C49.589 0 63.89 14.301 63.89 31.945z"/><g fill="#fff"><ellipse cx="20.844" cy="28.449" rx="6.493" ry="9.727"/><ellipse cx="42.813" cy="28.449" rx="6.494" ry="9.727"/></g><g fill="#25333a"><ellipse cx="20.844" cy="28.449" rx="4.201" ry="5.094"/><ellipse cx="42.812" cy="28.449" rx="4.201" ry="5.094"/></g>`),
		g.Group(children),
	)
}
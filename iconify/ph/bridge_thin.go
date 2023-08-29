package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 164h-36V88.09a67.81 67.81 0 0 0 34.5 31a4 4 0 1 0 3-7.42A59.77 59.77 0 0 1 196 56a4 4 0 0 0-8 0a60 60 0 0 1-120 0a4 4 0 0 0-8 0a59.77 59.77 0 0 1-37.5 55.64a4 4 0 0 0 3 7.42a67.81 67.81 0 0 0 34.5-31V164H24a4 4 0 0 0 0 8h36v28a4 4 0 0 0 8 0v-28h120v28a4 4 0 0 0 8 0v-28h36a4 4 0 0 0 0-8Zm-84-43v43h-40v-43a68 68 0 0 0 40 0ZM68 88a68.43 68.43 0 0 0 32 30v46H68Zm88 76v-46a68.43 68.43 0 0 0 32-30v76Z"/>`),
		g.Group(children),
	)
}
package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.29 4.31a.757.757 0 0 0-.82.16l-6.72 6.72V5c0-.3-.18-.58-.46-.69a.757.757 0 0 0-.82.16l-6.72 6.72V5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v14c0 .41.34.75.75.75s.75-.34.75-.75v-6.19l6.72 6.72a.75.75 0 0 0 1.28-.53v-6.19l6.72 6.72a.75.75 0 0 0 1.28-.53V5c0-.3-.18-.58-.46-.69Zm-9.04 12.88L6.06 12l5.19-5.19v10.38Zm8 0L14.06 12l5.19-5.19v10.38Z"/>`),
		g.Group(children),
	)
}
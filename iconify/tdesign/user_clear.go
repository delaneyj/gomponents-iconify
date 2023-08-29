package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm9.172 6.257L18 16.586l2.828-2.829l1.415 1.415L19.414 18l2.829 2.828l-1.415 1.415L18 19.414l-2.828 2.829l-1.415-1.415L16.586 18l-2.829-2.828l1.415-1.415ZM8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2H8Z"/>`),
		g.Group(children),
	)
}
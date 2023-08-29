package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telescope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m11.24 1l.59.24l2.11 4.93l-.23.59l-3.29 1.41l-.59-.24l-.17-.41L6.1 9l-.58-.19l-.16-.38L2.8 9.49l-.58-.24l-.72-1.67l.28-.59l2.5-1.06l-.18-.41l.24-.58L7.9 3.41L7.72 3L8 2.42L11.24 1zM2.5 7.64l.35.85l2.22-.91l-.37-.85l-2.2.91zm2.74-2.12l1.11 2.45l3-1.28l-1.11-2.44l-3 1.27zM8.79 3l1.86 4.11l2.29-1.01L11.18 2L8.72 3h.07zM8.5 9.1l3.02 4.9h-1.17l-1.88-3.03v4h-1V9.82L5.58 14h-1.1l1.7-3.9l2.32-1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
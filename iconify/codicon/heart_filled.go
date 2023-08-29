package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.88 4.78a3.489 3.489 0 0 0-.37-.9a3.24 3.24 0 0 0-.6-.79a3.78 3.78 0 0 0-1.21-.81a3.74 3.74 0 0 0-2.84 0a4 4 0 0 0-1.16.75l-.05.06l-.65.65l-.65-.65l-.05-.06a4 4 0 0 0-1.16-.75a3.74 3.74 0 0 0-2.84 0a3.78 3.78 0 0 0-1.21.81a3.55 3.55 0 0 0-.97 1.69a3.75 3.75 0 0 0-.12 1c0 .318.04.634.12.94a4 4 0 0 0 .36.89a3.8 3.8 0 0 0 .61.79L8 14.31l5.91-5.91c.237-.232.44-.498.6-.79A3.578 3.578 0 0 0 15 5.78a3.747 3.747 0 0 0-.12-1Z"/>`),
		g.Group(children),
	)
}
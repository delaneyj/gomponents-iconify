package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudXBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 36a92.09 92.09 0 0 0-81 48.36A68 68 0 1 0 72 220h88a92 92 0 0 0 0-184Zm0 160H72a44 44 0 0 1-1.82-88A91.86 91.86 0 0 0 68 128a12 12 0 0 0 24 0a68 68 0 1 1 68 68Zm32.49-75.51L177 136l15.52 15.51a12 12 0 0 1-17 17L160 153l-15.51 15.52a12 12 0 1 1-17-17L143 136l-15.52-15.51a12 12 0 1 1 17-17L160 119l15.51-15.52a12 12 0 0 1 17 17Z"/>`),
		g.Group(children),
	)
}
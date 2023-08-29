package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelicopterPad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 .055c-4.418 0-8 3.566-8 7.968c0 4.4 3.582 7.968 8 7.968s8-3.567 8-7.968C16 3.621 12.418.055 8 .055zm.004 15.057c-3.934 0-7.121-3.181-7.121-7.105C.883 4.083 4.071.902 8.004.902c3.933 0 7.121 3.181 7.121 7.105c0 3.924-3.187 7.105-7.121 7.105z"/><path d="M8.018 2.08c-3.264 0-5.91 2.654-5.91 5.927c0 3.273 2.646 5.927 5.91 5.927c3.264 0 5.911-2.654 5.911-5.927c0-3.273-2.648-5.927-5.911-5.927zm2.059 8.039h-1.14V9.062H7.062v1.057H5.944V5.961h1.118v1.914h1.875V5.961h1.14v4.158z"/></g>`),
		g.Group(children),
	)
}
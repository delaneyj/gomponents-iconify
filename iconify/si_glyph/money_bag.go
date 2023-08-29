package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9 11h.922v1.959H9zM7 8h.938v1.974H7z"/><path d="M9.296 4.779c.858-1 1.47-3.257 1.47-4.074c0-1.726-1.029.44-2.298.44c-1.271 0-2.298-2.188-2.298-.44c0 .831.629 3.148 1.512 4.123c-2.608.74-4.664 4.494-4.664 7.794c0 2.918 2.455 3.364 5.482 3.364s5.481-.486 5.481-3.364c0-3.006-2.004-7.201-4.685-7.843zm1.735 8.268h-1v.984h-1v1.031H7.969v-1.031h-1v-.984l.013-.016H5.969v-1.062h1.047v1.02l.016-.02h.938v-1.938H7v-.984H5.953V7.985H7V6.969h.969V5.938h1.062v1.031h1v1h.984v1.062H9.968V8.01l-.016.006H9.03v1.953h1.016v1.016h.984v2.062h.001z"/></g>`),
		g.Group(children),
	)
}
package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.822 12.062l-.864.007L9.987 11H7.03l-.01 1.06l-.902-.002c-1.19.809-2.097 2.134-2.097 2.983c0 1.229 8.929 1.25 8.929 0c0-.453-.827-2.049-2.128-2.979zm-1.779 2H7.958v-1.104h1.085v1.104zm2.814-12.039c.076-.512.122-.874.122-.992c0-1.312-6.917-1.291-6.917 0c0 1.311 1.682 8.901 3.031 8.901h.733c.448 0 .892-.84 1.287-2.005c3.229-.413 3.882-3.235 3.886-3.293V2.023h-2.142zm-1.378 4.956c.419-1.303.769-2.87.983-4h1.558l.001 1.692c-.009.085-.238 1.891-2.542 2.308z"/>`),
		g.Group(children),
	)
}
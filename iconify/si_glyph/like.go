package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Like(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.811 13.958H2.194a1.135 1.135 0 0 1-1.146-1.122V7.128c0-.62.515-1.123 1.146-1.123h2.617c.634 0 1.148.503 1.148 1.123v5.708c0 .62-.515 1.122-1.148 1.122zm10.127-3.009v1.084h1.514c-.076 1.146-.658 1.897-1.74 1.897h-4.426c-.688 0-1.029-1.312-2.699-1.312l-.558.019V6.681s1.063-.166 1.419-1.291c0 0 1.451-3.961 2.57-4.356c.658 0 1.191-.047 1.191 1.125l-.595 1.814s-.353 2.032 2.14 2.032h2.145c.688 0 1.06.424 1.06 1.049c0 0 .014.357.007.896h-2.027v1.084h1.99a16.57 16.57 0 0 1-.218 1.916h-1.773v-.001z"/>`),
		g.Group(children),
	)
}
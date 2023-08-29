package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alien(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.014.143c-3.855 0-6.983 2.979-6.983 6.651C2.031 10.469 7.209 16 9.014 16C10.822 16 16 10.469 16 6.794C16 3.122 12.873.143 9.014.143zM7.895 10.895c-.316.318-1.414-.271-2.448-1.321C4.411 8.528 3.829 7.42 4.145 7.1c.315-.321 1.412.269 2.447 1.317c1.033 1.048 1.619 2.155 1.303 2.478zm2.219-.008c-.32-.32.27-1.426 1.321-2.473c1.049-1.047 2.158-1.636 2.48-1.314c.32.318-.271 1.424-1.32 2.47c-1.05 1.047-2.161 1.634-2.481 1.317z"/>`),
		g.Group(children),
	)
}
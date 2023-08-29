package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mushrooms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.002 1C4.149 1 1.025 4.446 1.025 7.115c0 2.666 3.124 2.869 6.977 2.869c3.854 0 6.978-.203 6.978-2.869C14.979 4.446 11.855 1 8.002 1zm-.08 10.169c-1.047 0-2.032-.019-2.892-.115c-.001.087-.005.172-.005.259c0 3.142-3.579 5.688 2.945 5.688c6.525 0 2.947-2.546 2.947-5.688c.001-.092-.003-.182-.004-.273c-.884.107-1.903.129-2.991.129z"/>`),
		g.Group(children),
	)
}
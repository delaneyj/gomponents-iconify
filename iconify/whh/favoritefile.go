package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Favoritefile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.27 1024h-768q-26 0-45-18.5T.27 960V65q0-27 19-45.5t45-18.5h448v352q0 13 9.5 22.5t22.5 9.5h352v575q0 27-18.5 45.5t-45.5 18.5zm-435-404l-77-172l-77 172l-179 24l132 130l-34 186l158-91l158 91l-34-186l132-130zm179-620q26 0 44 18l257 257q19 19 19 46h-320V0z"/>`),
		g.Group(children),
	)
}
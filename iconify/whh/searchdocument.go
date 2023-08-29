package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Searchdocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.27 1024h-768q-26 0-45-18.5T.27 960V65q0-27 19-45.5t45-18.5h448v352q0 13 9.5 22.5t22.5 9.5h352v575q0 27-18.5 45.5t-45.5 18.5zm-261-185l-92-92q33-48 33-107q0-80-56-136t-135.5-56t-136 56t-56.5 136t56.5 136t135.5 56q59 0 107-33l92 92q5 5 12.5 5t13.5-5l26-26q5-6 5-13.5t-5-12.5zm-251-71q-53 0-90.5-37.5t-37.5-90.5t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5zm256-768q26 0 44 18l257 257q19 19 19 46h-320V0z"/>`),
		g.Group(children),
	)
}
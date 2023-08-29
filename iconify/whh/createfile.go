package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Createfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M831.973 1024h-767q-27 0-45.5-18.5T.973 960V65q0-26 19-45t45-19h448v352q0 13 9 22.5t23 9.5h351v574q0 27-18.5 45.5t-45.5 18.5v1zm-255-384h-128V512q0-26-19-44.5t-45.5-18.5t-45 18.5t-18.5 44.5v128h-128q-27 0-45.5 19t-18.5 45t18.5 44.5t45.5 18.5h128v128q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5V767h128q26 0 45-18.5t19-44.5t-19-45t-45-19zm0-640q21 0 44 19l256 257q19 19 19 46h-319V0z"/>`),
		g.Group(children),
	)
}
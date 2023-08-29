package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Locket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.681 15.064c-.99-3.928-4.304-6.972-8.308-7.584c-4.381-.67-8.327 1.448-10.373 4.84c-2.047-3.392-5.993-5.51-10.376-4.84c-4.004.613-7.316 3.658-8.306 7.586a10.148 10.148 0 0 0 .369 6.19c2.517 7.242 11.418 15.007 15.816 18.512a3.999 3.999 0 0 0 4.994 0c4.398-3.505 13.299-11.27 15.816-18.511a10.14 10.14 0 0 0 .368-6.193Z"/>`),
		g.Group(children),
	)
}
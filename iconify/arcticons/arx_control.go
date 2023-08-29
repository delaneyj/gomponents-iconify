package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArxControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.756 34.756v3.706a3.985 3.985 0 0 0 3.492 4.02a3.873 3.873 0 0 0 4.252-3.854V23.57a3.442 3.442 0 0 0-3.442-3.442H24.166a3.985 3.985 0 0 0-4.02 3.492A3.873 3.873 0 0 0 24 27.872h3.872a6.884 6.884 0 0 1 6.884 6.884Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.5a18.5 18.5 0 0 1 0-37h14.628a3.872 3.872 0 1 1 0 7.744H24a10.756 10.756 0 0 0 0 21.512h.86a3.872 3.872 0 1 1 0 7.744Z"/>`),
		g.Group(children),
	)
}
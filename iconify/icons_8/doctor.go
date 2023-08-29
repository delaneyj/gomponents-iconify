package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doctor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.438 6l-.282.47S5 11.652 5 20v1h6.094A4.923 4.923 0 0 0 11 22c0 2.75 2.25 5 5 5s5-2.25 5-5c0-.344-.027-.675-.094-1H27v-1c0-4.61-.776-7.99-1.563-10.22c-.786-2.228-1.625-3.374-1.625-3.374L23.5 6H8.437zm1.218 2h12.72c.145.208.573.732 1.186 2.47c.65 1.84 1.232 4.73 1.344 8.53H20c-.915-1.21-2.376-2-4-2s-3.085.79-4 2H7.094c.228-6.65 2.232-10.43 2.562-11zM15 10v2h-2v2h2v2h2v-2h2v-2h-2v-2h-2zm1 9c1.67 0 3 1.33 3 3s-1.33 3-3 3s-3-1.33-3-3s1.33-3 3-3zm0 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}
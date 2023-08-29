package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeGloves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 6.5a1.5 1.5 0 0 1 3 0V19h1V8.5a1.5 1.5 0 0 1 3 0V24l2.923-3.653a1.881 1.881 0 0 1 3.034 2.218l-4.25 6.374a9.116 9.116 0 0 1-3.616 3.15L30 43H18l.905-10.859A5.997 5.997 0 0 1 16 27V10.5a1.5 1.5 0 0 1 3 0V19h1V8.5a1.5 1.5 0 0 1 3 0V19h1V6.5ZM12 12h2v16c0 1.85 1.092 3.784 2.517 4.644l.544.328L16.216 41h-2.011l.738-7.011C13.181 32.647 12 30.286 12 28V12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
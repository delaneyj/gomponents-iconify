package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Traderepublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.755 19.266a.965.965 0 0 0 .745-.934h-.004V9.685A.97.97 0 0 0 42.3 8.75l-13.896 3.296a2.407 2.407 0 0 1-1.294-.05l-6.273-2.039a2.39 2.39 0 0 0-1.298-.05L5.245 13.308a.962.962 0 0 0-.745.932v8.648a.973.973 0 0 0 1.202.934l13.763-3.295a2.409 2.409 0 0 1 1.298.05l6.273 2.038a2.428 2.428 0 0 0 1.294.05l14.425-3.402Zm0 15.425a.96.96 0 0 0 .745-.934h-.004V25.11a.97.97 0 0 0-1.197-.934L28.403 27.47a2.407 2.407 0 0 1-1.294-.05l-6.273-2.038a2.39 2.39 0 0 0-1.298-.05L5.245 28.733a.96.96 0 0 0-.745.934v8.648a.973.973 0 0 0 1.202.934l13.763-3.295a2.409 2.409 0 0 1 1.298.05l6.273 2.038a2.4 2.4 0 0 0 1.294.05Z"/>`),
		g.Group(children),
	)
}
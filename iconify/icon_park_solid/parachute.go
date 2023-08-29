package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parachute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSParachute0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M16.727 24c2.424-2.424 4.849-3.636 7.273-3.636c2.424 0 4.849 1.212 7.273 3.636c2.828-2.424 4.95-3.636 6.363-3.636c1.414 0 3.536 1.212 6.364 3.636c0-11.046-8.954-20-20-20S4 12.954 4 24c2.828-2.424 4.95-3.636 6.364-3.636c1.414 0 3.535 1.212 6.363 3.636Z"/><path stroke-linecap="round" d="m4 24l20 20l-7.273-20m14.546 0L24 44l20-20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSParachute0)"/>`),
		g.Group(children),
	)
}
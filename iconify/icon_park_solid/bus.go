package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBus0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M6.012 34.005V8.036a2 2 0 0 1 2-2H40a2 2 0 0 1 2 2v25.969a3 3 0 0 1-3 3h-1.995V38a4 4 0 1 1-8 0v-.995h-9.997v.997a3.998 3.998 0 0 1-7.997 0v-.997H9.012a3 3 0 0 1-3-3Z"/><path stroke-linecap="round" d="M42 23H6"/><path fill="#fff" stroke-linecap="round" d="M34 13H14v10h20V13Z"/><path stroke-linecap="round" d="M14 30h2m16 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBus0)"/>`),
		g.Group(children),
	)
}
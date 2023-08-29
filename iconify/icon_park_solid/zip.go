package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSZip0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M4.629 12.749L24 5l19.371 7.749a1 1 0 0 1 .629.928V42a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V13.677a1 1 0 0 1 .629-.928Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M24 22v10"/><path stroke="#000" d="M33 27c1.657 0 3-1.12 3-2.5S34.657 22 33 22h-2.7a.3.3 0 0 0-.3.3v2.2c0 1.38 1.343 2.5 3 2.5Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M30 22v10"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M12 22h6.005L12 32h6.005"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSZip0)"/>`),
		g.Group(children),
	)
}
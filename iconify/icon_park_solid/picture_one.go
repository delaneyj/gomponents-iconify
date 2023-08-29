package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPictureOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path fill="#fff" d="M18 23a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path d="M42 36L31 26l-10 9l-7-6l-8 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPictureOne0)"/>`),
		g.Group(children),
	)
}
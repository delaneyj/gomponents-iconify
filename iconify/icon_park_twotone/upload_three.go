package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUploadThree0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" d="m32 28l-8-8l-8 8m8 8V20.5M16 14h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUploadThree0)"/>`),
		g.Group(children),
	)
}
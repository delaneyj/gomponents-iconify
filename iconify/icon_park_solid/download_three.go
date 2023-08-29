package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDownloadThree0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke="#000" stroke-linecap="round" d="m32 28l-8 8l-8-8m8-8v15.5M16 14h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDownloadThree0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUploadLaptop0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M25 9H11a3 3 0 0 0-3 3v21h32v-9"/><path fill="#fff" d="M4 33h40v2a6 6 0 0 1-6 6H10a6 6 0 0 1-6-6v-2Z"/><path stroke-linecap="round" d="M37 19V7m-5 5l5-5l5 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUploadLaptop0)"/>`),
		g.Group(children),
	)
}
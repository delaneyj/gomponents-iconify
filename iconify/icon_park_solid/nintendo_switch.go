package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NintendoSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNintendoSwitch0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M6 18c0-5.657 0-8.485 1.757-10.243C9.515 6 12.343 6 18 6h2v36h-2c-5.657 0-8.485 0-10.243-1.757C6 38.485 6 35.657 6 30V18Zm36 0c0-5.657 0-8.485-1.757-10.243C38.485 6 35.657 6 30 6h-2v36h2c5.657 0 8.485 0 10.243-1.757C42 38.485 42 35.657 42 30V18Z"/><path stroke="#000" d="M35 13v2M13 33v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNintendoSwitch0)"/>`),
		g.Group(children),
	)
}
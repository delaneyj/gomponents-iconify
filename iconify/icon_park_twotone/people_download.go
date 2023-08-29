package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPeopleDownload0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M19 20a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path stroke-linecap="round" d="M37 40V30m-4 6l4 4l4-4m-14-8h-8.2c-4.48 0-6.72 0-8.432.872a8 8 0 0 0-3.496 3.496C6 34.08 6 36.32 6 40.8V42h21"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPeopleDownload0)"/>`),
		g.Group(children),
	)
}
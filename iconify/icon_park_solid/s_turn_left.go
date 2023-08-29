package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSTurnLeft0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M42 38H13c-4 0-7-3-7-7s3-7 7-7h6m0 0h16c4 0 7-3 7-7s-3-7-7-7H17"/><path stroke-linecap="round" stroke-linejoin="round" d="m37 43l5-5l-5-5"/><circle cx="12" cy="10" r="5" fill="#fff" transform="rotate(-180 12 10)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSTurnLeft0)"/>`),
		g.Group(children),
	)
}
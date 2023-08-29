package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCopyOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M13 38h28V16H30V4H13v34Z"/><path stroke-linejoin="round" d="m30 4l11 12M7 20v24h21"/><path d="M19 20h4m-4 8h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCopyOne0)"/>`),
		g.Group(children),
	)
}
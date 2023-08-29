package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterKeyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEnterKeyOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M23 23V5h20v38H5V23h18Z"/><path d="M33 13v20H13"/><path d="m17 29l-4 4l4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEnterKeyOne0)"/>`),
		g.Group(children),
	)
}
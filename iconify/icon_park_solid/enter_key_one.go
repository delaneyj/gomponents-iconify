package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterKeyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEnterKeyOne0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M23 23V5h20v38H5V23h18Z"/><path stroke="#000" d="M33 13v20H13"/><path stroke="#000" d="m17 29l-4 4l4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEnterKeyOne0)"/>`),
		g.Group(children),
	)
}
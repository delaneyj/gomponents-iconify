package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transporter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTransporter0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M42 8H20a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2ZM4 34h14V20h-7l-7 6.462V34Z"/><path stroke-linecap="round" d="M18 36a4 4 0 0 1-8 0m30 0a4 4 0 0 1-8 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTransporter0)"/>`),
		g.Group(children),
	)
}
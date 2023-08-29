package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handheld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandheld0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M42 18v26H6V18"/><path fill="#555" d="M42 4H6v14h36V4Z"/><path stroke-linecap="round" d="M16 27v8m-4-4h8"/><path fill="#555" d="M32 35a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandheld0)"/>`),
		g.Group(children),
	)
}
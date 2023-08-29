package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrakePads(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBrakePads0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Z"/><path fill="#555" d="M24 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-24a20 20 0 0 1 20 20h-7.994A12.006 12.006 0 0 0 24 11.994V4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBrakePads0)"/>`),
		g.Group(children),
	)
}
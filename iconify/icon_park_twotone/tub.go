package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTub0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M40 23V12a7 7 0 0 0-7-7v0a7 7 0 0 0-7 7v1"/><path fill="#555" d="M40 29v-6H8v6a8 8 0 0 0 8 8h16a8 8 0 0 0 8-8Z"/><path d="M43 23H5"/><path stroke-linejoin="round" d="m17 37l-4 6m18-6l4 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTub0)"/>`),
		g.Group(children),
	)
}
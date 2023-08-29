package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FSixKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFSixKey0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path d="M30 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 18.876C33.382 17.19 31.892 16 30.152 16C27.86 16 26 17.451 26 20v8m-5-12h-7v16m0-8h7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFSixKey0)"/>`),
		g.Group(children),
	)
}
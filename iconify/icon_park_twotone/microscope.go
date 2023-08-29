package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMicroscope0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M26 44c4.346-3.008 6.663-6.049 6.95-9.122c.286-3.073-.642-5.366-2.784-6.878"/><path fill="#555" fill-rule="evenodd" d="M27.655 28.223a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z" clip-rule="evenodd"/><path stroke-linecap="round" stroke-linejoin="round" d="m24.288 27l-5.51 5.577l-9.192-9.192L27.97 5l9.192 9.192l-6.187 6.187"/><path stroke-linecap="round" d="m6.558 28.136l7.861 7.678M6 44h36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMicroscope0)"/>`),
		g.Group(children),
	)
}
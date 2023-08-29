package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boxing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBoxing0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" stroke-miterlimit="2" d="M38 36v6a2 2 0 0 1-2 2H17a2 2 0 0 1-2-2v-6"/><path fill="#fff" stroke-linejoin="round" stroke-miterlimit="2" d="M11 15v-5a6 6 0 0 1 6-6h21a6 6 0 0 1 6 6v20c0 3.314-2.68 6-5.994 6H15C8 36 4 28 4 23v-8h7Z"/><path d="M11 14v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBoxing0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDocSearch0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path fill="#555" d="M28 16c0 1.38-.56 2.63-1.465 3.535A5 5 0 1 1 28 16Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m30 23l-3.465-3.465m0 0a5 5 0 1 0-7.071-7.07a5 5 0 0 0 7.072 7.07ZM17 30h14m-14 6h7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDocSearch0)"/>`),
		g.Group(children),
	)
}
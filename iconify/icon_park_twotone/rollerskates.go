package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rollerskates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRollerskates0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M44 4H28v8h16V4Z"/><path fill="#555" d="M44 12v22c0 1.11-.89 2-2 2H8c-2.21 0-4-1.79-4-4v-6c0-4.42 3.58-8 8-8h16v-6h16Z"/><path d="M14 24v-6m7 6v-6m2 0H12M9 44a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm10.33 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm10.34 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM40 44a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRollerskates0)"/>`),
		g.Group(children),
	)
}
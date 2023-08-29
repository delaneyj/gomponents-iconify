package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peoples(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPeoples0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M19 20a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path d="M32.608 7A6.996 6.996 0 0 1 36 13a6.996 6.996 0 0 1-3.392 6"/><path fill="#555" d="M4 40.8V42h30v-1.2c0-4.48 0-6.72-.872-8.432a8 8 0 0 0-3.496-3.496C27.92 28 25.68 28 21.2 28h-4.4c-4.48 0-6.72 0-8.432.872a8 8 0 0 0-3.496 3.496C4 34.08 4 36.32 4 40.8Z"/><path d="M44 42v-1.2c0-4.48 0-6.72-.872-8.432a8 8 0 0 0-3.496-3.496"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPeoples0)"/>`),
		g.Group(children),
	)
}
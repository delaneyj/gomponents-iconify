package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafeNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 15.5h-9a2 2 0 0 1-2-2v-9h-18a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-26Zm-11-11l11 11"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.612 23.804h12.776c.484 0 .877.392.877.877v9.442a.877.877 0 0 1-.877.877H17.612a.877.877 0 0 1-.877-.877V24.68c0-.485.393-.878.877-.878h0Z"/><path d="M19.192 23.804V21.79a4.805 4.805 0 0 1 9.608 0v2.014"/><circle cx="23.996" cy="29.402" r="1.863"/></g>`),
		g.Group(children),
	)
}
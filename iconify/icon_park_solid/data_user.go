package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDataUser0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M22 8v12c0 2.21-4.03 4-9 4s-9-1.79-9-4V8"/><path d="M22 14c0 2.21-4.03 4-9 4s-9-1.79-9-4"/><path fill="#fff" d="M22 8c0 2.21-4.03 4-9 4s-9-1.79-9-4s4.03-4 9-4s9 1.79 9 4Z"/><path d="M32 6h6a4 4 0 0 1 4 4v6M16 42h-6a4 4 0 0 1-4-4v-6"/><circle cx="35" cy="29" r="5" fill="#fff"/><path fill="#fff" d="M44 44H26a9 9 0 1 1 18 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDataUser0)"/>`),
		g.Group(children),
	)
}
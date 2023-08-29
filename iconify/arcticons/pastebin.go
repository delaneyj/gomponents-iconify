package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pastebin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.37 10.35H7a2.48 2.48 0 0 0-2.5 2.48v22.34a2.5 2.5 0 0 0 5 0v-6.83h3.89c5.74 0 9-3.28 9-9s-3.28-8.99-9.02-8.99Zm0 13H9.48v-8h3.89c3 0 4.05 1.09 4.05 4s-1.1 4.01-4.05 4.01Zm21.1-3.69h-3.89v-.45a2.5 2.5 0 0 0-5 0v16a2.48 2.48 0 0 0 2.48 2.48h6.39c5.74 0 9-3.27 9-9s-3.24-9.03-8.98-9.03Zm0 13h-3.89v-8h3.89c3 0 4.05 1.08 4.05 4s-1.1 4.01-4.05 4.01Z"/><circle cx="28.08" cy="12.83" r="2.48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
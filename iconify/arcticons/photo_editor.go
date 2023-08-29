package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 24V10.5a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2H24"/><circle cx="14.158" cy="17.169" r="2.772" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.875 29.375L14.25 26l5.375 5.375l13.812-13.812l3.688 3.688m2.183 6.189c.959-.804 2.45-.648 3.254.31s.69 2.411-.269 3.215l-1.529-1.79l-1.456-1.735Zm-9.57 7.57l9.57-7.57l1.456 1.735l1.53 1.79l-9.545 7.636m-3.012-3.59l-1.54 4.232l4.552-.642"/>`),
		g.Group(children),
	)
}
package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyBill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.44 5.78H4.56a2.507 2.507 0 0 0-2.5 2.5v7.44a2.514 2.514 0 0 0 2.5 2.5h14.88a2.507 2.507 0 0 0 2.5-2.5V8.28a2.5 2.5 0 0 0-2.5-2.5ZM3.06 8.28a1.5 1.5 0 0 1 1.5-1.5h1.48a3.521 3.521 0 0 1-2.98 2.98Zm1.5 8.94a1.511 1.511 0 0 1-1.5-1.5v-1.48a3.521 3.521 0 0 1 2.98 2.98Zm16.38-1.5a1.5 1.5 0 0 1-1.5 1.5h-1.48a3.521 3.521 0 0 1 2.98-2.98Zm0-2.49a4.528 4.528 0 0 0-3.99 3.99h-9.9a4.528 4.528 0 0 0-3.99-3.99v-2.46a4.528 4.528 0 0 0 3.99-3.99h9.9a4.528 4.528 0 0 0 3.99 3.99Zm0-3.47a3.521 3.521 0 0 1-2.98-2.98h1.48a1.5 1.5 0 0 1 1.5 1.5Z"/><circle cx="12.002" cy="11.998" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlymDecsync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.205 10.214A16.646 16.646 0 0 0 7.24 24m1.313 0h2.891l-3.977 6.068L3.5 24h5.053zm30.894 0h-2.891l3.967-6.068L44.5 24h0h-5.053zM14.69 37.796A16.636 16.636 0 0 0 40.633 24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.856 13.379a.226.226 0 0 1 .186-.378c3.882.576 9.294 5.773 9.294 12.014c0 6.472-4.746 10.157-5.26 9.638c-.35-.355 2.467-2.257 1.58-3.166c-.519-.519-2.054 1.693-2.82.92c-.955-.93 1.985-5.58-.396-7.945c0 0-3.786 3.826-3.786 6.738s2.782 3.143 2.539 3.707s-6.529-1.394-6.529-7.133c0-5.185 5.959-10.642 5.959-12.555a2.664 2.664 0 0 0-.767-1.84Z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Somafm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.233 24.762H42.5v17.766H13.233z"/><circle cx="33.795" cy="33.645" r="5.215" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.822" cy="29.681" r="1.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.574" cy="29.681" r="1.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.822 38.86v-5.651m12.888-8.447L16.683 16.95M5.5 14.689l3.371 3.927l.515-3.49l3.332 4.481m6.341-14.136l-1.117 5.053l3.101-1.682l-1.585 5.354"/>`),
		g.Group(children),
	)
}
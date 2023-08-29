package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrayerTimes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.5 5.5h-21c-4.4 0-8 3.6-8 8v21c0 4.4 3.6 8 8 8h21c4.4 0 8-3.6 8-8v-21c0-4.4-3.6-8-8-8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.333 37C12.95 37 11 36.056 11 33.673V14c0-1.65 1.35-3 3-3h20c1.65 0 3 1.35 3 3v18.667A4.346 4.346 0 0 1 32.667 37v-1.526c-1.214-1.005-1.693-1.21-1.693-1.21s1.524-1.23 1.898-1.756c.816-1.15 2.863-6.045-1.469-9.095c-1.552-1.093-6.31-3.317-6.31-3.317c-1.036-2.482-1.128-3.545-1.093-6.172c.035 2.627-.057 3.69-1.094 6.172c0 0-4.756 2.224-6.309 3.317c-4.332 3.05-2.285 7.945-1.469 9.095c.374.527 1.898 1.756 1.898 1.756s-.479.205-1.692 1.21V37Z"/>`),
		g.Group(children),
	)
}
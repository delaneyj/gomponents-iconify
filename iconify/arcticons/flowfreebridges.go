package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flowfreebridges(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="11.57" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.57" cy="36.43" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.78" cy="36.43" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.78" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.43" cy="36.35" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.21" cy="24" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.43" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.21" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="36.43" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="30.11" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.57 34.36V13.64m18.64 4.12v-4.12M24 17.78h-6.22m12.43 0h-2.8m-9.63 16.58v-16.6m8.29 18.59h4.14M19.86 11.57H24m6.21 18.54v6.24m6.22-22.71v16.47m-6.22 0h6.22M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM24 15.92v-4.35m0 16.47v-8.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19.63a1.86 1.86 0 0 0 0-3.71"/>`),
		g.Group(children),
	)
}
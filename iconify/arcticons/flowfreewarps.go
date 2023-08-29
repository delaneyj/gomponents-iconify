package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flowfreewarps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="17.78" cy="30.11" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.57" cy="30.11" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.57" cy="17.76" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.78" cy="17.76" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.43" cy="30.11" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.43" cy="17.76" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.21" cy="36.43" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.43" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="30.11" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.21" cy="30.11" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.21 17.76v-4.12m-18.64-2.07H24M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM24 28.04V17.76M24 5.5v6.07m0 30.93v-6.07m-12.43 0H23.8m12.63-4.25v4.25m-4.15 0h4.15m6.07-12.59h-6.07m0-4.01v4.01m-6.22 4.2V11.57m4.15 0h-4.15M11.57 32.18v4.25M5.5 23.84h12.19m-6.12-8.15v-4.12m8.28 6.19H24m-6.22 10.28v-4.2"/>`),
		g.Group(children),
	)
}
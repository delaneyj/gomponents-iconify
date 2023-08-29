package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flowfree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="11.57" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.78" cy="36.43" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.78" cy="30.2" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="30.2" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.43" cy="30.2" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.43" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.21" cy="11.57" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.21" cy="17.78" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="36.43" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="23.96" r="2.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.57 36.35V13.64M24 21.89v-4.11m4.14 0H24m4.14-6.21H17.78m0 16.56V11.57m-6.21 24.78h4.14m10.36.08h10.36m0-4.16v4.16M26.07 30.2h4.14m0-6.24v6.24m6.22-16.56v10.32m-6.22 0h6.22M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
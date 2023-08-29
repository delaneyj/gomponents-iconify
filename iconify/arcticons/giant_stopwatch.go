package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiantStopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 26.087a4.157 4.157 0 0 0 4.174 4.174a4.028 4.028 0 0 0 4.02-4.174v-4.174a4.126 4.126 0 0 0-4.02-4.174A4.258 4.258 0 0 0 7.5 21.913Zm7.265-6.802l-6.183 9.43m12.937-2.628a4.157 4.157 0 0 0 4.174 4.174a4.028 4.028 0 0 0 4.019-4.174v-4.174a4.126 4.126 0 0 0-4.02-4.174a4.258 4.258 0 0 0-4.173 4.174Zm7.265-6.802l-6.183 9.43m9.705-2.628a4.157 4.157 0 0 0 4.175 4.174a4.028 4.028 0 0 0 4.019-4.174v-4.174a4.126 4.126 0 0 0-4.02-4.174a4.258 4.258 0 0 0-4.173 4.174Zm7.267-6.802l-6.184 9.43"/><circle cx="18.702" cy="27.811" r=".75" fill="currentColor"/><circle cx="18.702" cy="20.189" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}
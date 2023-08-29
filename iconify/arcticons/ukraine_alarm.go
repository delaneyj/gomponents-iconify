package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UkraineAlarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.173 20.124h4.878v7.752h-4.878z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.051 27.876c7.697 0 12.533 3.364 17.748 8.075V12.05c-5.215 4.71-10.051 8.075-17.748 8.075M33.8 27.028a3.028 3.028 0 1 0 0-6.056m-22.627 6.904v13.379"/>`),
		g.Group(children),
	)
}
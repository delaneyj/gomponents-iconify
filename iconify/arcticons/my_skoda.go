package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MySkoda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.71 30.87c.841.924 3.112 3.952 3.533 5.382h2.691l2.23-1.43h16.584v1.683s4.822-1.01 5.79-2.818c-1.178-2.06-3.873-3-5.383-3.238v1.906h-2.482s2.075-2.972 2.075-9.813s-2.916-11.551-7.178-15.7c-9.42.448-16.426 3.64-18.729 14.971c7.626.897 11.874 6.252 11.089 10.738a10.766 10.766 0 0 0-4.08-1.745Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.185 8.35c3.086 1.66 6.338 5.865 6.338 5.865m-13.773 0a20.305 20.305 0 0 1 6.848 3.56"/><circle cx="27.028" cy="28.093" r="2.047" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
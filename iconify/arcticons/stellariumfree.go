package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stellariumfree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.16 36.79c3.87-2.7 12.57-4.64 17.1-3.56m-8.51-21.91a6.9 6.9 0 0 1-5.37 12.25a6.9 6.9 0 1 0 5.37-12.25Z"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.79 18.05h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.52 9.84h0"/><circle cx="17.01" cy="11.68" r=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.71" cy="25.08" r=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.12" cy="22.9" r="1" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="10.8" cy="11.91" r="1" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.2 44.17l-11.32-19.6l-3.58 6.2l-1.66-2.88L12 30.73L5.54 19.57l-2.71 4.69"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.69 19.13h0M16.83 17h0"/>`),
		g.Group(children),
	)
}
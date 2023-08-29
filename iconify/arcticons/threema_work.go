package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreemaWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5c-10.08 0-18.26 6.32-18.26 14.11a12.66 12.66 0 0 0 5.06 9.73L7.34 35l10.2-3.21a22.88 22.88 0 0 0 6.46.93c10.08 0 18.26-6.31 18.26-14.1S34.08 4.5 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8.65a5 5 0 0 1 5 5v4.15h.71a1 1 0 0 1 .95.95h0V26a1 1 0 0 1-.95.95h-11.4a1 1 0 0 1-.95-.95h0v-7.27a1 1 0 0 1 .95-.95H19v-4.15a5 5 0 0 1 5-4.98Z"/><circle cx="13.63" cy="40.6" r="2.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="40.6" r="2.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.37" cy="40.6" r="2.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.02 17.78h9.96m-2.41 1.965l-1.28 5.12l-1.28-5.12l-1.28 5.12l-1.28-5.12"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MujAlbert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.55 16.82c.894-2.402.046-4.285-1.837-4.85c-2.73-.847-7.155 2.26-10.027 7.203a17.608 17.608 0 0 0-2.071 5.555a8.035 8.035 0 0 0 1.506 6.214a7.483 7.483 0 0 1-.424-.941c-1.13-2.966 2.73-8.05 10.357-12.005"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 30.048c.047 3.342-8.238 6.12-18.548 6.12c-10.31.047-18.69-2.636-18.737-6.026c-.047-3.342 8.286-6.073 18.549-6.12c10.31-.047 18.689 2.636 18.736 6.026Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.965 24.116c4.519-3.295 6.779-6.779 5.366-8.709c-1.93-2.542-9.65-1.365-17.277 2.59S4.568 27.034 5.698 30a6.78 6.78 0 0 0 1.6 2.118a3.007 3.007 0 0 1-1.083-1.977"/>`),
		g.Group(children),
	)
}
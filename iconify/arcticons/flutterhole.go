package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flutterhole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="19.5" ry="10.636"/><ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="14.891" ry="6.736"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.377 29.568a8.417 8.417 0 0 0-16.755 0M24 13.364v8.599m0 12.673v-3.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.9 23.542c.863-2.861 2.875-6.971 4.675-8.809m-1.772 12.491c1.06-2.176 6.002-7.434 9.358-8.279m-4.801 8.813c2.309-1.499 5.42-1.494 6.699-1.496"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.31 23.85a10.556 10.556 0 0 1 6.58 4.206m-5.984 2.13a9.954 9.954 0 0 1 4.99 2.636m-15.797-9.28c-.862-2.861-2.874-6.971-4.675-8.809m1.773 12.491c-1.06-2.176-6.002-7.434-9.358-8.279m4.801 8.813c-2.309-1.499-5.42-1.494-6.699-1.496"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.69 23.85a10.556 10.556 0 0 0-6.58 4.206m5.984 2.13a9.954 9.954 0 0 0-4.99 2.636"/>`),
		g.Group(children),
	)
}
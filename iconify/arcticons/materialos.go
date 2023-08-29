package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Materialos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.962 24a9.96 9.96 0 0 0-17.584-6.411l-.004-.003a4.147 4.147 0 0 1-3.263 1.328a6.302 6.302 0 0 1-5.89-8.54a15.83 15.83 0 0 0 23.39 21.074a44.9 44.9 0 0 0 .324-.292l-.004-.004A9.93 9.93 0 0 0 33.96 24Z"/><circle cx="24" cy="24" r="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.057 7.981A6.303 6.303 0 0 1 21.77 2.5a15.828 15.828 0 0 0-13.3 14.375"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.125 8.47a6.303 6.303 0 0 1 6.501-1.249a15.83 15.83 0 0 0-25.092 11.666"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.019 18.057A6.303 6.303 0 0 1 45.5 21.77a15.834 15.834 0 0 0-30.396-3.251"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.58 31.18a6.303 6.303 0 0 1 1.2 6.446a15.83 15.83 0 0 0-23.392-21.074"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.943 40.02a6.303 6.303 0 0 1-3.714 5.48a15.83 15.83 0 0 0-1.637-31.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.875 39.53a6.303 6.303 0 0 1-6.501 1.249a15.83 15.83 0 0 0 21.074-23.39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.98 29.943A6.302 6.302 0 0 1 2.5 26.23a15.83 15.83 0 0 0 31.442-1.638"/>`),
		g.Group(children),
	)
}
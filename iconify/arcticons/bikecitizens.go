package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bikecitizens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.274 7.75h-13l9.165 19.655M35.274 7.75v3.245"/><circle cx="32.125" cy="28.875" r="1.625" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.125 17.5A11.375 11.375 0 1 1 20.75 28.875m-16.25 0a8.125 8.125 0 0 1 16.25 0m-5.688-11.349a4.875 4.875 0 1 1 2.42-4.647"/>`),
		g.Group(children),
	)
}
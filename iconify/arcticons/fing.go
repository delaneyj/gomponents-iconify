package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="6" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="42" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.272" cy="36.728" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.272" cy="11.272" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.728" cy="36.728" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="6" cy="24" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.484 41.918a18.147 18.147 0 0 0 6.721-2.786m-20.41 0a18.147 18.147 0 0 0 6.72 2.786M6.085 27.499a18.149 18.149 0 0 0 2.783 6.706M8.86 13.807a18.148 18.148 0 0 0-2.775 6.694M20.5 6.085a18.148 18.148 0 0 0-6.693 2.775m25.807 5.687A18.27 18.27 0 0 0 27.5 6.085m11.632 28.12A18.165 18.165 0 0 0 42.25 24a18.455 18.455 0 0 0-.119-2.092"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.165 9.577l-.551 4.97l-4.969-.551"/>`),
		g.Group(children),
	)
}
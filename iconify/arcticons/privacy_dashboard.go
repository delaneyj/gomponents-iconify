package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacyDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.753 18.688a4.634 4.634 0 0 1 9.269 0m0 0v2.789m-9.269-2.789v2.752"/><circle cx="22.387" cy="27.343" r="1.613" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.387 30.534v-1.528"/><rect width="12.848" height="12.848" x="15.963" y="21.098" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.545"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".845" d="M42.447 34.748A21.499 21.499 0 1 1 23.828 2.5"/>`),
		g.Group(children),
	)
}
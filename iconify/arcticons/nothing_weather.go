package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NothingWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="10.245" cy="11.467" r="3.745" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.262 32.488c5.654 0 10.238-4.583 10.238-10.237s-4.583-10.238-10.238-10.238c-4.596 0-8.485 3.03-9.779 7.202a7.39 7.39 0 0 0-4.578-1.586a7.43 7.43 0 0 0 0 14.86h14.357Z"/><circle cx="31.806" cy="36.096" r=".836" fill="currentColor"/><circle cx="28.461" cy="39.441" r=".836" fill="currentColor"/><circle cx="24" cy="36.096" r=".836" fill="currentColor"/><circle cx="20.655" cy="39.441" r=".836" fill="currentColor"/><circle cx="16.194" cy="36.096" r=".836" fill="currentColor"/><circle cx="12.849" cy="39.441" r=".836" fill="currentColor"/>`),
		g.Group(children),
	)
}
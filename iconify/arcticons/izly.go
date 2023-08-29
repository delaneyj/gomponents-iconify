package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Izly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.721 25.503v4.642a3.439 3.439 0 0 1-3.438 3.439h0a3.428 3.428 0 0 1-2.432-1.008"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.721 19.828v5.674a3.439 3.439 0 0 1-3.438 3.44h0a3.439 3.439 0 0 1-3.44-3.44v-5.674m-15.765 0h6.877l-6.877 9.113h6.877"/><circle cx="11.478" cy="15.616" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.478 19.828v9.113m14.107-13.755v12.036a1.72 1.72 0 0 0 1.72 1.72h.516m.024-14.115a2.518 2.518 0 0 0-2.269-2.281l-.006.006m4.595 2.126a4.95 4.95 0 0 0-4.456-4.455m6.821 4.23a7.32 7.32 0 0 0-6.603-6.603"/>`),
		g.Group(children),
	)
}
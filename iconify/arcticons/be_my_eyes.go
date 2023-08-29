package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeMyEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.28 30.72l-7.61 7.61m12.81-4.95L20.8 44.07m12.46-1.99l-4.96-9.6m13.78.78l-9.6-4.96m11.59-7.5l-10.69 1.7m4.95-12.83l-7.61 7.61M27.2 3.93l-1.7 10.69m-5.8.9l-4.96-9.6m.78 13.78l-9.6-4.96m8.7 10.76L3.93 27.2M14.5 24h-12m13.78 5.54l-9.7 7.14m14.52-3.64L17.38 44.4m13.24 0L26.9 33.04m14.52 3.64l-9.7-7.14M45.5 24h-12m7.92-12.57l-9.73 6.98M30.62 3.6L26.9 14.96m-5.8 0L17.38 3.6m-1.07 14.81l-9.73-6.98M24 45.5v-12m12.68 7.92l-7.14-9.7m14.86-1.1L33.04 26.9m11.36-9.52L33.04 21.1m3.64-14.52l-7.09 9.73M24 14.5v-12m-5.59 13.81l-6.98-9.73m3.53 14.52L3.6 17.38m11.36 9.52L3.6 30.62m14.81 1.07l-6.98 9.73m5.85-24.14L11.1 11.1m3.52 11.4l-8.7-1.37m9.6 7.17l-7.84 3.97m12.02.21l-3.97 7.84m11.14 1.76l-1.35-8.7M36.9 36.9l-6.18-6.18m11.36-3.85l-8.7-1.37m6.94-9.77l-7.84 3.97m-.21-12.02l-3.97 7.84m-5.8-.9l-1.37-8.7"/><circle cx="24" cy="24" r="9.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
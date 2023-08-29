package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pojavlauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.86 14.3L23.89 4.5L7.03 14.23L24 24.03l16.86-9.73z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 24.03l-16.97-9.8V33.7L24 43.5V24.03zm16.86-9.73L24 24.03V43.5l16.86-9.73V14.3z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.08" d="m24 32.28l-3.48-6.53l-10-5.77l-3.49 2.5m33.83.05l-3.45-2.48l-9.95 5.74L24 32.28m10.26-18.03L23.91 8.28l-10.28 5.94l10.35 5.97l10.28-5.94zm-6.9-3.98l-10.28 5.94m13.73-3.95L20.53 18.2m10.3-1.97l-10.34-5.97m6.92 7.95l-10.35-5.97"/>`),
		g.Group(children),
	)
}
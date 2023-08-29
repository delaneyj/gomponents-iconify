package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mapsme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.66 22.15l4.39-4.3l2.45 13.36M32.74 10.8l2-.34a4.77 4.77 0 0 1 5.59 3.78l.73 3.61M20.21 23l7.65-7.58L43.5 31.21m-21.35 0a1.9 1.9 0 0 1-1.06.83h0a1.7 1.7 0 0 1-.75-3.32h0a1.7 1.7 0 0 1 2 1.21a1.84 1.84 0 0 1-.15 1.28Zm-5-17.42l-2.15.39c-.66-4.35 1.11-6.18 3.8-6.67h0l3.2-.62c2.69-.5 5.36.85 6.08 4.8l-2.3.44M12 14.79l20.76-4m-24.4 4.64l3.64-.64l15.91 16l-6.63 6.59l-16.77-17a4.75 4.75 0 0 1 3.85-4.95ZM7.42 35.8L4.51 20.38m18.65 18.89l-9.85 1.84a4.81 4.81 0 0 1-5.59-3.78l-.3-1.53l6.21-6.19m14.26 1.13l6.27 6.33l-7.65 1.47l-3.35.73l-1.9-1.94m22.24-6.12A4.68 4.68 0 0 1 39.69 36h0l-5.5 1.1l-6.27-6.33l8.77-8.59m-16.48.83l-6.58 6.6m-.27-13.44L6.84 22.7m21.02-7.27l-3.08-3.1"/>`),
		g.Group(children),
	)
}
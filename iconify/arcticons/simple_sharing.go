package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimpleSharing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.438 23.281a11.338 11.338 0 0 0-22.665 0a6.206 6.206 0 0 0 .15 12.41h22.365a6.206 6.206 0 0 0 .15-12.41Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.754 32.98l4.217-4.217l-4.217-4.217"/><circle cx="12.099" cy="28.763" r="1.007" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="5.507" cy="28.763" r="1.007" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.93" cy="28.763" r="1.007" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
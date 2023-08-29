package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleGalleryGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="30" x="5.5" y="9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35 9v23.5H12v-17l23 .101M42.466 32.5H35"/><circle cx="29" cy="21.601" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 21.601H5.5M22.9 32.5c0-6.02-4.88-10.9-10.9-10.9"/>`),
		g.Group(children),
	)
}
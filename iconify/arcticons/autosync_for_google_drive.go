package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutosyncForGoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="11.638" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.194 2.682c4.032 3.402 4.409 8.07 4.409 9.79M44.7 29.83c-4.865 1.149-8.883.4-10.372-.46M8.286 38.674c.783-5.05 3.52-9.242 5.009-10.101"/>`),
		g.Group(children),
	)
}
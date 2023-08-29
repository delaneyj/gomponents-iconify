package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobilepass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c13.355-4.243 16.635-14.368 16.635-19.286V11.309a1.133 1.133 0 0 0-.739-1.062L24.788 4.64a2.266 2.266 0 0 0-1.576 0L8.104 10.247a1.133 1.133 0 0 0-.74 1.063v12.905c0 4.918 3.281 15.043 16.636 19.286Z"/><circle cx="24.989" cy="23.54" r="1.904" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.773" cy="23.54" r="1.904" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.556" cy="23.54" r="1.904" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.302 22.419l2.513 2.513l5.533-5.532"/>`),
		g.Group(children),
	)
}
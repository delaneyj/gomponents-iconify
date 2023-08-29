package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.502 2.135A1 1 0 0 1 18 3v4a3.99 3.99 0 0 1 2.981 1.333A3.989 3.989 0 0 1 22 11c0 1.024-.386 1.96-1.019 2.667A3.993 3.993 0 0 1 18 15v4a1 1 0 0 1-1.496.868L10 16.152V21a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h5.734l6.77-3.868a1 1 0 0 1 .998.003zM10 14a1 1 0 0 1 .496.132L16 17.277V4.723l-5.504 3.145A1 1 0 0 1 10 8H4v6h6zm-4 2v4h2v-4H6zm12-3c.592 0 1.123-.256 1.491-.667c.317-.354.509-.82.509-1.333s-.192-.979-.509-1.333A1.993 1.993 0 0 0 18 9v4z"/>`),
		g.Group(children),
	)
}
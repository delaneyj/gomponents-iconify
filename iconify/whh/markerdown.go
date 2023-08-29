package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markerdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m507 748l-123 276l-123-276Q146 709 73 609T0 384q0-104 51.5-192.5t140-140T384 0t192.5 51.5t140 140T768 384q0 125-73 225T507 748zM384 128q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75z"/>`),
		g.Group(children),
	)
}
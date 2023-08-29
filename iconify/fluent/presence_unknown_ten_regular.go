package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceUnknownTenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.999 1a3.999 3.999 0 1 0 0 7.997a3.999 3.999 0 0 0 0-7.997ZM0 4.999a4.999 4.999 0 1 1 9.997 0a4.999 4.999 0 0 1-9.997 0Z"/>`),
		g.Group(children),
	)
}
package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v1zm-8.775 5.11c.243-.163.379-.409.408-.738V5.72l2.432.898a1 1 0 0 0 1.285-.591c.191-.518-.121-1.066-.639-1.258A381.257 381.257 0 0 1 9.946 3.41c-.441-.163-.755-.005-1.03.204c-.182.14-.291.408-.326.806l.043 1.49l-2.428-.877a1 1 0 1 0-.68 1.881c1.656.606 2.912 1.065 3.77 1.377c.34.124.566.06.93-.183z"/>`),
		g.Group(children),
	)
}
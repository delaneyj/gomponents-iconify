package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quoteup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 448H768v128q0 27 18.5 45.5t45 18.5t45.5 19t19 45.5t-19 45t-45 18.5H704q-31 0-61-32t-48.5-77t-18.5-83V128q0-53 37.5-90.5T704 0h192q53 0 90.5 37.5T1024 128v192q0 53-37.5 90.5T896 448zm-576 0H192v128q0 27 18.5 45.5T256 640t45.5 19t18.5 45.5t-18.5 45T256 768H128q-30 0-60.5-32t-49-77T0 576V128q0-53 37.5-90.5T128 0h192q53 0 90.5 37.5T448 128v192q0 53-37.5 90.5T320 448z"/>`),
		g.Group(children),
	)
}
package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.002 1c0-.553.444-1 .993-1h3.972c.549 0 .993.447.993 1v14c0 .553-.444 1-.993 1H1.995a.996.996 0 0 1-.993-1V1zm8 0c0-.553.444-1 .993-1h3.972c.549 0 .993.447.993 1v14c0 .553-.444 1-.993 1H9.995a.996.996 0 0 1-.993-1V1z"/>`),
		g.Group(children),
	)
}
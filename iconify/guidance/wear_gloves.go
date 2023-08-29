package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WearGloves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5 15.5v-3a3 3 0 0 0-3-3h-.5V20A3.5 3.5 0 0 1 5 23.5h4M12 2V.5h-.5a3 3 0 0 0-3 3v4m.17-5H8a3 3 0 0 0-3 3V16m10.5-4V3.5a3 3 0 0 1 3-3h.5V12V5.5a3 3 0 0 1 3-3h.5v21H12A3.5 3.5 0 0 0 8.5 20V9.5H9a3 3 0 0 1 3 3v3m3.67-13H15a3 3 0 0 0-3 3m0 0V12m0-6.5V16"/>`),
		g.Group(children),
	)
}
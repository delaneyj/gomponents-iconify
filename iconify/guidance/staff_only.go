package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StaffOnly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M19.5 20.5v.5m0-.5A3.5 3.5 0 0 0 16 24m3.5-3.5v-11H19a3 3 0 0 0-3 3v3m0 0a3.5 3.5 0 0 0-3.5 3.5v1m3.5-4.5v-13h-.5a3 3 0 0 0-3 3V12V.5H12a3 3 0 0 0-3 3V12V2.5h-.5a3 3 0 0 0-3 3V24"/>`),
		g.Group(children),
	)
}
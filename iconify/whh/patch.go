package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M962.716 962q-62 62-150 62t-150-62l-84-85q-10-40 19-70l210-209q30-31 72-19l83 83q62 62 62 150t-62 150zm-416-215q-21 21-50 21t-50-21l-169-169q-21-21-21-50.5t21-49.5l201-201q21-21 50-21t50 21l169 169q21 20 21 49.5t-21 50.5zm-330-320q-29 29-70 19l-84-84q-62-62-62-150t62-150t150-62t150 62l83 82q12 42-19 73z"/>`),
		g.Group(children),
	)
}
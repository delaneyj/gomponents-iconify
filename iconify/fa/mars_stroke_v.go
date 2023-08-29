package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarsStrokeV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M640 644q217 24 364.5 187.5T1152 1216q0 167-87 306t-236 212t-319 54q-133-15-245.5-88t-182-188T2 1263q-12-155 52.5-292t186-224T512 644V512H352q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h160V219l-92 92q-10 9-23 9t-22-9l-46-46q-9-9-9-22t9-23L531 19q19-19 45-19t45 19l202 201q9 10 9 23t-9 22l-46 46q-9 9-22 9t-23-9l-92-92v165h160q14 0 23 9t9 23v64q0 14-9 23t-23 9H640v132zm-64 1020q185 0 316.5-131.5T1024 1216T892.5 899.5T576 768T259.5 899.5T128 1216t131.5 316.5T576 1664z"/>`),
		g.Group(children),
	)
}
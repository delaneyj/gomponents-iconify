package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="309.912" cy="82.088" r="34.088" fill="currentColor"/><path fill="currentColor" d="m322 143.462l-60.585-20.64l-77.615 28.226A24.073 24.073 0 0 0 168 173.6V232h32v-52.793l48.811-17.749L158.622 440h33.613l50.082-155.811l7.871 2.568L288 356.079V440h32v-85.96a24.068 24.068 0 0 0-2.931-11.493l-37.56-68.861l28.949-88.715l19.27 34.684A24.011 24.011 0 0 0 348.707 232H400v-32h-46.586Z"/>`),
		g.Group(children),
	)
}
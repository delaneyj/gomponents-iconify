package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenDesktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M40 392h200v72h-80v32h192v-32h-80v-72h200a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24H40a24.028 24.028 0 0 0-24 24v296a24.028 24.028 0 0 0 24 24Zm8-312h416v280H48Z"/>`),
		g.Group(children),
	)
}
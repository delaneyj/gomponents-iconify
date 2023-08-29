package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeLetter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M400 163.2V68a20.023 20.023 0 0 0-20-20H132a20.023 20.023 0 0 0-20 20v95.2l-96 68.566V496h480V231.766Zm53.679 77.667L400 275.96v-73.44ZM144 80h224v216.883l-57.166 37.378l-46.578-24.152l-50.764 24.507L144 292.425Zm119.744 265.89L464 449.727V464H48v-13.957ZM48 271.575L179.144 351.2L48 414.509Zm295.446 79.6L464 272.347v141.334ZM112 202.52V273l-53.334-32.385Z"/>`),
		g.Group(children),
	)
}
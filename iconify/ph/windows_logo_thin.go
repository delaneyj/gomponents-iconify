package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 140h-80a4 4 0 0 0-4 4v57.45a4 4 0 0 0 3.28 3.94l80 14.55a4.37 4.37 0 0 0 .72.06a4 4 0 0 0 2.56-.93A4 4 0 0 0 220 216v-72a4 4 0 0 0-4-4Zm-4 71.21l-72-13.09V148h72ZM104 140H40a4 4 0 0 0-4 4v40a4 4 0 0 0 3.28 3.94l64 11.63a3.51 3.51 0 0 0 .72.07a4 4 0 0 0 4-4V144a4 4 0 0 0-4-4Zm-4 50.84l-56-10.18V148h56ZM218.56 36.93a4 4 0 0 0-3.28-.87l-80 14.55a4 4 0 0 0-3.28 3.94V112a4 4 0 0 0 4 4h80a4 4 0 0 0 4-4V40a4 4 0 0 0-1.44-3.07ZM212 108h-72V57.88l72-13.09ZM103.28 56.43l-64 11.63A4 4 0 0 0 36 72v40a4 4 0 0 0 4 4h64a4 4 0 0 0 4-4V60.36a4 4 0 0 0-4.72-3.93ZM100 108H44V75.34l56-10.18Z"/>`),
		g.Group(children),
	)
}
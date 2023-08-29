package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftOutlookLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 96a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm0 48a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm128-40h-8V48a16 16 0 0 0-16-16h-80a16 16 0 0 0-16 16v16H40a16 16 0 0 0-16 16v96a16 16 0 0 0 16 16h32v16a16 16 0 0 0 16 16h128a16 16 0 0 0 16-16v-88a16 16 0 0 0-16-16ZM112 48h80v77.24l-40 28.89V80a16 16 0 0 0-16-16h-24ZM40 176V80h96v96H40Zm48 32v-16h48a16 16 0 0 0 16-16v-2.13L199.26 208Zm128-7.65L165.66 164L216 127.65Z"/>`),
		g.Group(children),
	)
}
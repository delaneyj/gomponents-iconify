package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h6.5v2H3V1Zm12 2.414V7h3.586L15 3.414ZM18 14.5c-.69 0-1.25.56-1.25 1.25v.75h5.749V23h-9v-6.5h1.251v-.75a3.25 3.25 0 0 1 5.541-2.305l.71.705l-1.41 1.418l-.71-.705A1.243 1.243 0 0 0 18 14.5Zm-2.501 4V21h5v-2.5h-5Z"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWaterfallsV0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M23 20V6H6v14h17Zm19 22V28H25v14h17ZM31 6v14h11V6H31ZM6 28v14h11V28H6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWaterfallsV0)"/>`),
		g.Group(children),
	)
}
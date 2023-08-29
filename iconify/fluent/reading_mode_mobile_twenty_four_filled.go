package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadingModeMobileTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 2.001a2.25 2.25 0 0 1 2.245 2.096L20 4.25v15.498a2.25 2.25 0 0 1-2.096 2.245l-.154.005H6.25a2.25 2.25 0 0 1-2.245-2.096L4 19.75V4.251a2.25 2.25 0 0 1 2.096-2.245l.154-.005h11.5Zm-5.502 10.996H7.75l-.102.007a.75.75 0 0 0 0 1.487l.102.006h4.498l.102-.006a.75.75 0 0 0 0-1.487l-.102-.007ZM16.25 10h-8.5l-.102.006a.75.75 0 0 0 0 1.487l.102.007h8.5l.102-.007a.75.75 0 0 0 0-1.487L16.25 10Zm0-2.999h-8.5l-.102.007a.75.75 0 0 0 0 1.486l.102.007h8.5l.102-.007a.75.75 0 0 0 0-1.486L16.25 7Z"/>`),
		g.Group(children),
	)
}
package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMaldives(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#699635" d="M2 32c0 9.8 4.7 18.5 12 24h36c7.3-5.5 12-14.2 12-24S57.3 13.5 50 8H14C6.7 13.5 2 22.2 2 32z"/><path fill="#ed4c5c" d="M14 8h36c-5-3.8-11.2-6-18-6S19 4.2 14 8m18 54c6.8 0 13-2.2 18-6H14c5 3.8 11.2 6 18 6"/><path fill="#fff" d="M43 49.6C34.6 48 28.3 40.7 28.3 32S34.6 16 43 14.4c-1.2-.2-2.4-.4-3.7-.4C29.2 14 21 22.1 21 32s8.2 18 18.3 18c1.3 0 2.5-.1 3.7-.4"/>`),
		g.Group(children),
	)
}
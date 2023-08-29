package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForThailand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M52.4 10C47 5 39.9 2 32 2s-15 3-20.4 8h40.8zM11.6 54c5.4 5 12.5 8 20.4 8s15-3 20.4-8H11.6"/><path fill="#2a5f9e" d="M2 32c0 4.3.9 8.3 2.5 12h55c1.6-3.7 2.5-7.7 2.5-12s-.9-8.3-2.5-12h-55C2.9 23.7 2 27.7 2 32"/><path fill="#f9f9f9" d="M11.6 54h40.7c3-2.8 5.5-6.2 7.1-10h-55c1.8 3.8 4.2 7.2 7.2 10m40.8-44H11.6c-3 2.8-5.5 6.2-7.1 10h55c-1.7-3.8-4.1-7.2-7.1-10"/>`),
		g.Group(children),
	)
}
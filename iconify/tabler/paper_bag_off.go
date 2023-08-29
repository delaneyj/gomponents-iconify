package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperBagOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.158 3.185C7.414 3.066 7.7 3 8 3h8a2 2 0 0 1 2 2v1.82a5 5 0 0 0 .528 2.236l.944 1.888A5 5 0 0 1 20 13.18V16m-.177 3.824A2 2 0 0 1 18 21H6a2 2 0 0 1-2-2v-5.82a5 5 0 0 1 .528-2.236L6 8V6"/><path d="M13.185 13.173a2 2 0 1 0 2.64 2.647M6 21a2 2 0 0 0 2-2v-5.82a5 5 0 0 0-.528-2.236L6 8m5-1h2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
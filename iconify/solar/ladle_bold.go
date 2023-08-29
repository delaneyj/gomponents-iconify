package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5.684 2.75A2.934 2.934 0 0 0 2.75 5.684a.75.75 0 0 1-1.5 0a4.434 4.434 0 1 1 8.868 0v6.057a5.68 5.68 0 0 0-.974.618a3.7 3.7 0 0 0-.526.515v-7.19A2.934 2.934 0 0 0 5.684 2.75Zm2.949 13.393a7.066 7.066 0 0 0 14.084.225a3.763 3.763 0 0 1-.207.207c-.503.464-1.145.817-1.802 1.083c-1.324.536-3.02.842-4.708.842c-1.685 0-3.486-.305-4.908-.823c-.706-.256-1.4-.595-1.948-1.036a3.711 3.711 0 0 1-.51-.498Z"/><path d="M22 14.5c0 1.38-2.946 2.5-6 2.5s-6.5-1.12-6.5-2.5S12.946 12 16 12s6 1.12 6 2.5Z"/></g>`),
		g.Group(children),
	)
}
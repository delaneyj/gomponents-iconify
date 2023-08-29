package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerLogoYoutubeYoutubeClipSocialVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10" x=".5" y="2" rx="2"/><path d="M5.31 9.32v-4.2a.39.39 0 0 1 .6-.34l3.6 2.1a.4.4 0 0 1 0 .69l-3.6 2.1a.4.4 0 0 1-.6-.35Z"/></g>`),
		g.Group(children),
	)
}
package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerWebcamWebcamCameraFutureTechChatSkypeTechnologyVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="6.52" r="6"/><path d="M10.07 11.68a7.4 7.4 0 0 1 2.48 1.8m-11.1 0a7.4 7.4 0 0 1 2.48-1.8"/><circle cx="7" cy="6.52" r="3"/><circle cx="7" cy="6.52" r=".5"/></g>`),
		g.Group(children),
	)
}
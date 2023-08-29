package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerLogoWindowsTwoOsSystemMicrosoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 5.82h6.5v-4.5a.49.49 0 0 0-.56-.49L7 1.54Zm-2 0V1.71l-4.06.5a.5.5 0 0 0-.44.5v3.11Zm0 2H.5v3a.51.51 0 0 0 .42.5L5 12Zm2 0v4.4l5.92.95a.5.5 0 0 0 .58-.49V7.82Z"/>`),
		g.Group(children),
	)
}
package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerLogoWindowsOneOsSystemMicrosoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.94 2l12-1.38a.49.49 0 0 1 .56.49v11.82a.5.5 0 0 1-.58.49l-12-1.84a.51.51 0 0 1-.42-.5V2.46A.5.5 0 0 1 .94 2ZM6 1.38v10.98m7.5-5.6H.5"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.067 30l-.001-12l-6.44 8.061h7.95m-17.152-4.086a3.979 3.979 0 0 1 4.764-3.9c1.667.321 2.998 1.763 3.164 3.452c.123 1.257-.274 2.497-1.142 3.259C20.602 26.197 15.424 30 15.424 30h7.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.864 6.258C6.212 10.132 2.5 16.631 2.5 24c0 11.874 9.626 21.5 21.5 21.5S45.5 35.874 45.5 24S35.874 2.5 24 2.5a21.41 21.41 0 0 0-7.917 1.524"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VivaldiSnapshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.028 38.863l-.793 1.384a4.816 4.816 0 0 1-3.975 2.63a4.607 4.607 0 0 1-4.571-2.432c-3.048-5.253-6.062-10.527-9.089-15.792l-5.52-9.611a4.856 4.856 0 0 1 3.945-7.36a4.653 4.653 0 0 1 4.465 2.508c1.359 2.331 2.693 4.676 4.038 7.014c.97 1.684 1.92 3.38 2.917 5.047a7.626 7.626 0 0 0 6.391 3.953a7.784 7.784 0 0 0 8.19-6.92c.036-.32.06-.638.074-.8a8.156 8.156 0 0 0-.825-3.636a4.855 4.855 0 1 1 9.073-3.025a5.081 5.081 0 0 1-.665 3.281m0 0l-7.17 12.473"/><circle cx="34.5" cy="34.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.008 37.623a2.238 2.238 0 0 0 1.962.877h1.184a1.998 1.998 0 0 0 1.996-2h0a1.998 1.998 0 0 0-1.996-2h-1.308a1.998 1.998 0 0 1-1.996-2h0a1.998 1.998 0 0 1 1.996-2h1.184a2.238 2.238 0 0 1 1.962.876"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snackvideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.125 29.166A21.551 21.551 0 0 1 2.5 24C2.5 12.126 12.126 2.5 24 2.5S45.5 12.126 45.5 24S35.874 45.5 24 45.5c-7.116 0-13.425-3.457-17.338-8.784"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.506 11.5c-2.814 4.21.634 6.678.97 8.696S.16 30.963 2.5 35.28s9.575-1.85 11.594.505s2.55 5.608 5.691 5.58s5.285-.809 7.57 1.29s4.036 1.806 6.06.68"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.407 24.384a2.12 2.12 0 0 0 .002-3.133a32.661 32.661 0 0 0-5.743-4.202a33.264 33.264 0 0 0-6.52-2.873a2.118 2.118 0 0 0-2.704 1.566a32.66 32.66 0 0 0-.767 7.076a33.27 33.27 0 0 0 .772 7.085a2.115 2.115 0 0 0 2.697 1.56a32.661 32.661 0 0 0 6.522-2.877a33.268 33.268 0 0 0 5.74-4.202Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.698 26.397c.523-.477.524-1.3 0-1.778a18.53 18.53 0 0 0-3.258-2.384a18.878 18.878 0 0 0-3.699-1.63a1.202 1.202 0 0 0-1.534.888a18.537 18.537 0 0 0-.436 4.015c.002 1.546.194 2.909.438 4.02a1.2 1.2 0 0 0 1.53.886a18.54 18.54 0 0 0 3.701-1.633a18.873 18.873 0 0 0 3.258-2.384Z"/>`),
		g.Group(children),
	)
}
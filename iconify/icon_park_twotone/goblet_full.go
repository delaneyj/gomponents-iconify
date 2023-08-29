package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GobletFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGobletFull0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M33 44H13m10-16v16"/><path fill="#555" d="M35 16c0 1.675-.357 3.284-1 4.75C32.148 24.97 27.92 28 23 28a12 12 0 0 1-10.79-6.742A11.953 11.953 0 0 1 11 16c0-2.373.533-4.613 1.21-6.5C13.387 6.217 15 4 15 4h16s1.815 2.496 3 6.112c.574 1.752 1 3.767 1 5.888Z"/><path d="M35 16s-7 3-12 0s-12 0-12 0"/><path d="M34 10.112c.574 1.752 1 3.767 1 5.888c0 1.675-.357 3.284-1 4.75M12.21 9.5C11.533 11.387 11 13.627 11 16c0 1.886.435 3.67 1.21 5.258"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGobletFull0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBread0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M4 32.083c0-1.202.266-2.395.971-3.368C7.045 25.85 12.671 20 24 20c11.33 0 16.955 5.851 19.029 8.715c.705.973.971 2.166.971 3.368A7.917 7.917 0 0 1 36.083 40H11.917A7.917 7.917 0 0 1 4 32.083Z"/><path stroke="#fff" d="M12 9v4"/><path stroke="#000" d="M14 22v4"/><path stroke="#fff" d="M36 9v4"/><path stroke="#000" d="M34 22v4"/><path stroke="#fff" d="M24 7v6"/><path stroke="#000" d="M24 20v8"/><path stroke="#fff" d="M40 25.443C36.906 22.78 31.808 20 24 20s-12.906 2.779-16 5.443"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBread0)"/>`),
		g.Group(children),
	)
}
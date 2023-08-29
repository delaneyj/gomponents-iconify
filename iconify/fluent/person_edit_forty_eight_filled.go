package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonEditFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 4c-5.523 0-10 4.477-10 10s4.477 10 10 10s10-4.477 10-10S27.523 4 22 4ZM10.25 28A4.25 4.25 0 0 0 6 32.249V33c0 3.755 1.942 6.567 4.92 8.38C13.85 43.163 17.785 44 21.998 44c.002-.328.044-.664.132-1.003l.838-3.235a5 5 0 0 1 1.277-2.253L33.608 28H10.25Zm34.584-2.832a3.981 3.981 0 0 0-5.652.022L25.671 38.913a3 3 0 0 0-.767 1.351l-.838 3.234c-.383 1.477.961 2.82 2.437 2.438l3.235-.839a3 3 0 0 0 1.351-.766L44.812 30.82a3.981 3.981 0 0 0 .022-5.651Z"/>`),
		g.Group(children),
	)
}
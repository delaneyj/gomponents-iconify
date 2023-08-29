package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRadiation0"><g fill="none"><path fill="#fff" d="M24.004 27a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M19.04 31.476A8.94 8.94 0 0 0 24 32.969a8.94 8.94 0 0 0 4.96-1.493l6.061 9.207A19.867 19.867 0 0 1 24 44c-4.073 0-7.861-1.22-11.021-3.317l6.062-9.207Zm-4.024-6.992l-10.98.661A20.062 20.062 0 0 1 15.053 6l4.92 9.869a9.028 9.028 0 0 0-4.958 8.615Zm13.01-8.615L32.946 6a20.062 20.062 0 0 1 11.019 19.145l-10.98-.661a9.028 9.028 0 0 0-4.958-8.615Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRadiation0)"/>`),
		g.Group(children),
	)
}
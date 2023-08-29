package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M160 0c13.3 0 24 10.7 24 24v40h144V24c0-13.3 10.7-24 24-24s24 10.7 24 24v40h40c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H96c-35.3 0-64-28.7-64-64V128c0-35.3 28.7-64 64-64h40V24c0-13.3 10.7-24 24-24zm272 192H80v256c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16V192zM328 352H184c-13.3 0-24-10.7-24-24s10.7-24 24-24h144c13.3 0 24 10.7 24 24s-10.7 24-24 24z"/>`),
		g.Group(children),
	)
}
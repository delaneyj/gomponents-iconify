package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" d="M14.731 36.275s.178 3.763 3.763 3.763H24"/><path fill="none" stroke="currentColor" d="M33.26 43.41V15.647a3.464 3.464 0 0 0-1.732-3l-5.796-3.346a3.464 3.464 0 0 0-3.464 0l-5.796 3.346a3.464 3.464 0 0 0-1.732 3v27.854"/><path fill="none" stroke="currentColor" d="M33.269 36.275s-.178 3.763-3.763 3.763H24"/><rect width="13.227" height="3.063" x="17.386" y="29.142" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".738"/><rect width="3.381" height="2.56" x="17.231" y="34.041" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".783"/><rect width="3.381" height="2.56" x="22.309" y="34.041" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".783"/><rect width="3.381" height="2.56" x="27.388" y="34.041" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".783"/>`),
		g.Group(children),
	)
}
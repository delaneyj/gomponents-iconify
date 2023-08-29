package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headwear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHeadwear0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12.417 43C10.095 40.068 8 35.779 8 31c0-8.837 7.163-16 16-16s16 7.163 16 16c0 4.779-2.095 9.068-4.417 12"/><path fill="#555" d="M34 13.5L43 5l-3 12l-5 1l-1-4.5Zm-20 0L5 5l3 12l5 1l1-4.5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHeadwear0)"/>`),
		g.Group(children),
	)
}
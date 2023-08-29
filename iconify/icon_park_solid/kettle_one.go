package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KettleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSKettleOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M13 14h17s8 4.148 8 13.8c0 9.65-8 14.2-8 14.2H13s-7-6.022-7-14c0-7.979 7-14 7-14Z"/><path d="M38 28c-13-3-19 6-32 0m25-14h9s4 4 4 16"/><path fill="#fff" d="m9 6l23 2.667L31 14H13L9 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSKettleOne0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChafingDishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChafingDishOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 18h36v6a6 6 0 0 1-6 6H12a6 6 0 0 1-6-6v-6Z"/><path d="M40 42H8m5 0l3-12m19 12l-3-12m-2-12L27 4h-6l-3 14m18-7h4M8 11h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChafingDishOne0)"/>`),
		g.Group(children),
	)
}
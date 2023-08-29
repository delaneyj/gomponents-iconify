package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCrab0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 17C5 12 6 2 17 6c.946 3.782-2 11-12 11Zm38 0c0-5-1-15-12-11c-.946 3.782 2 11 12 11Z"/><rect width="32" height="20" x="8" y="22" fill="#555" rx="10"/><path d="M27 22v-4m-6 4v-4m22-1c1 3 1 9-3 12M5 17c-1 3 0 8 3 12m22 13h14l-4-9m-22 9H4l4-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCrab0)"/>`),
		g.Group(children),
	)
}
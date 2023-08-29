package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartIncreasing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M56.999 7.001h-5.218L37.01 25.902L27.959 15.41L2 53.059V62h5.982l20.862-30.255l8.547 9.908l19.608-25.09z"/><path fill="currentColor" d="M62 2H2v47.536L4.783 45.5H4V33h9.402l1.379-2H4V18.5h12.5v10.005l2-2.9V18.5h4.899l1.379-2H18.5V4H31v11.876l2.266 2.624h6.99l1.563-2H33V4h12.5v7.791l2-2.561V4H60v12.5h-.999v.751l-.977 1.249H60V31H48.256L45.5 34.525V45.5H33v-5.874l-2-2.317V45.5h-9.21l-1.379 2H31V60H18.5v-9.731l-2 2.9V60h-4.709l-1.379 2H62V2zM16.5 16.5H4V4h12.5v12.5zm29 43.5H33V47.5h12.5V60zM60 60H47.5V47.5H60V60zm0-14.5H47.5V33H60v12.5z"/>`),
		g.Group(children),
	)
}
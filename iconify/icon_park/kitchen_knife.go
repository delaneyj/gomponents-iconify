package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenKnife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 30C14 30 14 34 14 40C14 46 22 46 22 40C22 34 22 30 22 30"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 37H14"/><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M14 6C14 4.89543 14.8954 4 16 4H32.6349C32.9537 4 33.2672 4.07486 33.5228 4.26546C34.5417 5.0254 37.1977 7.58171 38 14C38.7728 20.1822 36.6313 26.3644 35.6182 28.8547C35.3295 29.5645 34.6334 30 33.8672 30H14V6Z"/><circle cx="22" cy="10" r="2" fill="#fff"/></g>`),
		g.Group(children),
	)
}
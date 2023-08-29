package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VegetableBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 14H7.37783C5.58056 14 4.18661 15.5695 4.39881 17.3542L7.11944 40.2361C7.23907 41.2423 8.09222 42 9.10545 42H39.4734C40.5033 42 41.3648 41.2179 41.4641 40.1927L43.6816 17.2891C43.8522 15.5262 42.4666 14 40.6955 14H38"/><path d="M5 22L43 22"/><path d="M15 29H33"/><path d="M17 36H31"/><path fill="#2F88FF" d="M24 6C19.5817 6 16 12.9247 16 21.4667C16 21.6452 16.0016 21.823 16.0047 22H31.9953C31.9984 21.823 32 21.6452 32 21.4667C32 12.9247 28.4183 6 24 6Z"/></g>`),
		g.Group(children),
	)
}
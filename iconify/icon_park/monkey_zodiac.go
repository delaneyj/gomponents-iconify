package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonkeyZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21.5925 18.1345C22.9329 16.6704 24.5079 13.1567 20.0846 10.8141C19.0793 10.1634 18.5 8.50002 17.5714 5.44585C13.7178 3.33107 5 3.00018 4 14.0004V44"/><path d="M17 25C21.5818 26.673 28.0333 30.2543 32 41C33.5333 44.7543 40.5067 45.8705 43.5 40C44.4978 38.0432 44 34.5035 39.5818 32.4503C36.4615 31.0002 34 26 38.5 24.0003C40.3478 23.397 42.4291 24.08 44 27.0002"/><path d="M29 35.0004C24.6552 33.8942 15.7724 34.5187 15 44"/></g>`),
		g.Group(children),
	)
}
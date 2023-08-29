package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#2F88FF" d="M24 4C27.9556 4 31.8224 5.17298 35.1114 7.37061C38.4004 9.56824 40.9638 12.6918 42.4776 16.3463C43.9913 20.0009 44.3874 24.0222 43.6157 27.9018C42.844 31.7814 40.9392 35.3451 38.1421 38.1421C35.3451 40.9392 31.7814 42.844 27.9018 43.6157C24.0222 44.3874 20.0009 43.9913 16.3463 42.4776C12.6918 40.9638 9.56824 38.4004 7.37061 35.1114C5.17298 31.8224 4 27.9556 4 24H24V4Z"/></g>`),
		g.Group(children),
	)
}
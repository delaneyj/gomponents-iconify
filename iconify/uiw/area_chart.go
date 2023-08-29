package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.087 2.054c.198 0 .357.16.357.357c0 .018-.007.034-.01.052h.01v7.755c0 .197.16.357.357.357h7.78c.196 0 .356.16.356.357l-.002.01l.002.085c0 4.955-4.015 8.973-8.969 8.973C4.015 20 0 15.982 0 11.027C0 6.07 4.015 2.054 8.968 2.054l.096.002c.008.002.015-.002.023-.002ZM11.13 0c4.877.05 8.82 3.996 8.871 8.876c0 .197-.16.357-.357.357c-.018 0-.034-.008-.052-.01v.01H11.13a.357.357 0 0 1-.357-.357V.41h.01c-.002-.018-.01-.034-.01-.052c0-.197.16-.357.357-.357Z"/>`),
		g.Group(children),
	)
}
package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m374.906 253.877l11.36-10.328l-109.911-120.9h-68.492L163.585 151.4A60.364 60.364 0 0 0 136 202.2V274h32v-71.8a28.477 28.477 0 0 1 13.013-23.967L208 160.712v93.836l105.7 220.381h32.942L272.714 315.7V166.214l74.628 82.086l82.409 226.626h25.537ZM214.7 96.861a34.081 34.081 0 1 0-10.871-24.949A33.96 33.96 0 0 0 214.7 96.861Z"/><path fill="currentColor" d="M132.796 474.929h34.465l57.514-143.785l-18.793-39.181l-73.186 182.966z"/>`),
		g.Group(children),
	)
}
package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xtend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M7.01 97.812h64.095L197.29 252.038L53.078 416.28H0v-15.022L124.183 254.04L7.01 110.831V97.81zm164.722-50.406H90.54v11.576l158.787 190.306L81.29 452.091v17.01h76.97l184.224-219.813L171.732 47.406zm334.012 14.353V40.728h-81.12L313.455 168.922l55.08 58.086l137.209-165.25zM362.635 273.715h-8.995l-39.197 40.686v11.068l121.91 145.803H512v-22.294L362.635 273.715z"/>`),
		g.Group(children),
	)
}
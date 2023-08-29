package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M800 434v434H655V434h145zm398 0v434h-145V434h145zm0 760l253-254V145H257v1049h326v217l217-217h398zM1596 0v1013l-434 434H836l-217 217H402v-217H4V289L113 0h1483z"/>`),
		g.Group(children),
	)
}
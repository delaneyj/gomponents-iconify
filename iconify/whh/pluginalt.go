package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pluginalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1025H256q-27 0-45.5-18.5T192 961v-1h352q40 0 68-28t28-68V544q0-39-28-67.5T544 448H192V65q0-26 18.5-45T256 1h384v352q0 13 9.5 22.5T672 385h351l1 1v575q0 27-18.5 45.5T960 1025zM704 0q26 0 44 19l257 257q19 18 18 45H704V0zM448 576v256q0 27-19 45.5T384 896H64q-27 0-45.5-18.5T0 832V576q0-26 18.5-45T64 512h320q27 0 45.5 19t18.5 45zm-69 117L214 580q-6-5-14-3t-8 7v56H96q-13 0-22.5 9.5T64 672v64q0 14 9.5 23t22.5 9h96v57q0 5 8 6.5t14-2.5l165-113q5-5 5-11.5t-5-11.5z"/>`),
		g.Group(children),
	)
}
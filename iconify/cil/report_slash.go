package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReportSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 352h32v32h-32zM355.411 16H156.589l-34.936 37.771l22.645 22.645L170.58 48h170.84L464 180.53v150.94l-31.063 33.585l22.644 22.645L496 344V168L355.411 16z"/><path fill="currentColor" d="M240 128v44.118l32 32V128h-32zM16 16v22.627l62.164 62.164L16 168v176l140.589 152h198.822l56.681-61.281L473.373 496H496v-22.627L38.627 16Zm325.42 448H170.58L48 331.47V180.53l52.808-57.095L240 262.627V320h32v-25.373l117.447 117.447Z"/>`),
		g.Group(children),
	)
}
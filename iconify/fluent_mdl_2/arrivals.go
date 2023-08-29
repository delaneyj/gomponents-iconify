package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrivals(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1216q0 40-15 75t-41 61t-61 41t-75 15H128L0 896h384l32 128h256L512 384h320l320 640h576q40 0 75 15t61 41t41 61t15 75zm-1692 64h1500q26 0 45-19t19-45q0-26-19-45t-45-19h-655q-78-162-158-320T753 512h-77q40 161 79 320t81 320H316l-16-64l-16-64H164l64 256zm1820 384v128H640v-128h1408z"/>`),
		g.Group(children),
	)
}
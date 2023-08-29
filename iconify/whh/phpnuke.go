package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phpnuke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 640h-64v128L384 576v-64h384q53 0 90.5-37.5T896 384t-37.5-90.5T768 256h-64V128h64q106 0 181 75t75 181t-75 181t-181 75zM256 256q-53 0-90.5 37.5T128 384t37.5 90.5T256 512h64v128h-64q-106 0-181-75T0 384t75-181t181-75h64V0l320 192v64H256z"/>`),
		g.Group(children),
	)
}
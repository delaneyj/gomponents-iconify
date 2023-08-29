package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delicious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M303 322H0V19h303v303zm352 0H352V19h303v303zM272 50H31v242h241V50zm31 624H0V371h303v303zm49 0V371h303v303H352zm47-272h-16v16zm41 0l-57 57v41l98-98h-41zm82 0L383 541v41l180-180h-41zM383 644h19l223-222v-20h-21L383 623v21zm101 0l141-141v-40L443 644h41zm82 0l59-59v-41L525 644h41zm41 0h18v-18z"/>`),
		g.Group(children),
	)
}
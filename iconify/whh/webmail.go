package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m667 512l276-276q81 127 81 276t-81 276zM119 184q72-86 174.5-135T512 0t218.5 49T905 184L512 577zM81 788Q0 661 0 512t81-276l276 276zm321-231l71 72q6 5 15.5 8t16.5 3h7q26 1 39-11l71-72l283 283q-72 86-174.5 135T512 1024t-218.5-49T119 840z"/>`),
		g.Group(children),
	)
}
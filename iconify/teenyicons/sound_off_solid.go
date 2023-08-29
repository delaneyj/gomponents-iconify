package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5c-.829 0-1.5.67-1.5 1.5V9.5c0 .829.67 1.5 1.5 1.5h1.862l4.88 2.929A.5.5 0 0 0 9 13.5v-12Zm4.207 5.996l1.414 1.413l-.707.707L12.5 8.203l-1.414 1.413l-.707-.707l1.414-1.413l-1.414-1.413l.707-.707L12.5 6.789l1.415-1.413l.706.707l-1.414 1.413Z"/>`),
		g.Group(children),
	)
}
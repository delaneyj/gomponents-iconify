package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M344 24Q323 3 289 3t-52 21L3 259v170h170l235-234q20-20 20.5-53T408 88zm-145 98l40 41l-108 108v-29l-41-9zm-43 265H62l20-20l-30-29l-7 6v-68l9-9l34 9v51l49 15l17 34l26-13zm34-34l-23-47l-6-3l110-110l41 40zm188-188l-38 38L229 92l38-38q9-9 22-9q11 0 25 9l64 64q7 7 7 23q0 17-7 24z"/>`),
		g.Group(children),
	)
}
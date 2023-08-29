package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M24 24v472h294.627L496 318.627V24Zm32 32h408v216H272v192H56Zm249.373 408H304V304h160v1.373Z"/><path fill="currentColor" d="M208 288v-80h80v-32h-80V96h-32v80H96v32h80v80h32z"/>`),
		g.Group(children),
	)
}
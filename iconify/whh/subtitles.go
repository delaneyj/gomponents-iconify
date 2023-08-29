package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subtitles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 768H128q-53 0-90.5-37.5T0 640V128q0-53 37.5-90.5T128 0h768q53 0 90.5 37.5T1024 128v512q0 53-37.5 90.5T896 768zM255.5 256q26.5 0 45.5 19t19 45h128q0-80-56-136t-136-56t-136 56t-56 136v128q0 80 56 136t136 56t136-56t56-136H320q0 27-19 45.5T255.5 512t-45-18.5T192 448V320q0-26 18.5-45t45-19zm512 0q26.5 0 45.5 19t19 45h128q0-80-56-136t-136-56t-136 56t-56 136v128q0 80 56 136t136 56t136-56t56-136H832q0 27-19 45.5T767.5 512t-45-18.5T704 448V320q0-26 18.5-45t45-19z"/>`),
		g.Group(children),
	)
}
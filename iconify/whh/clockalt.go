package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clockalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm256-576H576V192q0-27-19-45.5T511.5 128t-45 18.5T448 192v320q0 27 18.5 45.5T512 576h256q26 0 45-19t19-45.5t-18.5-45T768 448z"/>`),
		g.Group(children),
	)
}
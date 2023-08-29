package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mutealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1014 57L57 1014q-10 10-24 10t-23.5-9.5T0 991t10-24L967 10q10-10 24-10t23.5 9.5T1024 33t-10 24zM448 256L693 13q30-30 75 16v54L448 402V256zM128 704V320q0-26 19-45t45-19h192v210L131 720q-3-8-3-16zm640 293q-45 45-75 14L448 768v-18l320-320v567z"/>`),
		g.Group(children),
	)
}
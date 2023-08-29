package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conversation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M185 448h325v163q0 19-13 32t-34 14H185L69 750v-93H46q-19 0-32-14T0 611V309q0-19 14-33t32-13h139v185zM695 8q20 0 33 13t13 34v301q0 19-13 32t-33 14h-23v93l-116-93H232V55q0-20 13-34t33-13h417z"/>`),
		g.Group(children),
	)
}
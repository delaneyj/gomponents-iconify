package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Good(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M505 572v37h109v70H243l-59-45V315s82-61 118-136c36-76 34-164 34-164h94s4 72 0 129c-5 58-24 129-24 129h260v82H505v37h161v73H505v36h130v71H505zM0 629h142V321H0v308z"/>`),
		g.Group(children),
	)
}
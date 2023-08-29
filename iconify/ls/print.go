package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m181 57l-11 86h377l-11-86c0-26-21-46-47-46H227c-26 0-46 20-46 46zM46 190h625c26 0 46 20 46 46v274c0 26-20 46-46 46h-85v127H131V556H46c-26 0-46-20-46-46V236c0-26 20-46 46-46zm494 447V340H177v297h363zm-54-211H226c-6 0-11-6-11-12v-14c0-6 5-11 11-11h260c6 0 12 5 12 11v14c0 6-6 12-12 12zm0 83H226c-6 0-11-6-11-12v-14c0-6 5-11 11-11h260c6 0 12 5 12 11v14c0 6-6 12-12 12zm0 83H226c-6 0-11-6-11-12v-14c0-6 5-11 11-11h260c6 0 12 5 12 11v14c0 6-6 12-12 12z"/>`),
		g.Group(children),
	)
}
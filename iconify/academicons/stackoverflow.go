package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stackoverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="m252.083 64l-27.489 20.396l101.987 137.46l27.488-20.396Zm-63.404 54.986l-21.73 26.155l131.693 109.53l21.729-26.164zm-50.552 70.947l-14.186 31.481l155.197 72.272l14.186-31.039zm-29.264 77.599l-7.093 33.698l167.608 35.032l7.093-33.707zm-78.932 43.45V448h308.177V310.981h-34.14v102.88H64.079V310.98Zm68.288 34.148v34.14h171.159v-34.14z"/>`),
		g.Group(children),
	)
}
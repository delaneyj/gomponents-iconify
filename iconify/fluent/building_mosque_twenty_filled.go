package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingMosqueTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.25 4.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm14.25.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm-7.296-1.456a.5.5 0 0 0-.408 0h-.002l-.003.002l-.012.005l-.04.02l-.151.07a14.624 14.624 0 0 0-2.13 1.271c-.584.423-1.187.939-1.648 1.528C5.35 7.527 5 8.225 5 9c0 .711.148 1.387.416 2H3V6.5a.5.5 0 0 0-1 0V17a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1a2 2 0 1 1 4 0v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V6.5a.5.5 0 0 0-1 0V11h-2.416A4.983 4.983 0 0 0 15 9c0-.775-.35-1.473-.81-2.06c-.461-.589-1.064-1.105-1.647-1.527a14.621 14.621 0 0 0-2.281-1.343l-.041-.019l-.012-.005l-.003-.002h-.002Z"/>`),
		g.Group(children),
	)
}
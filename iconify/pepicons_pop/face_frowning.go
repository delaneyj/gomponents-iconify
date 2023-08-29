package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceFrowning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 17.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm0 2a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Z" clip-rule="evenodd"/><path d="M8.5 7.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M12.221 14.627A1 1 0 0 0 13.8 13.4l-.8.6l.8-.601l-.002-.001l-.002-.003l-.003-.005l-.01-.011a1.891 1.891 0 0 0-.08-.095a2.538 2.538 0 0 0-.191-.188a3.597 3.597 0 0 0-.698-.478C12.192 12.286 11.28 12 10 12s-2.192.286-2.814.618a3.6 3.6 0 0 0-.698.478a2.557 2.557 0 0 0-.272.283l-.009.011l-.003.005l-.002.003H6.2c0 .002 0 .002.8.602l-.8-.6a1 1 0 0 0 1.579 1.227a1.598 1.598 0 0 1 .348-.245C8.442 14.214 9.029 14 10 14c.97 0 1.558.214 1.873.382a1.6 1.6 0 0 1 .348.245Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10.5 7.5a2.5 2.5 0 0 1 5 0v4a2.5 2.5 0 0 1-5 0v-4Z" clip-rule="evenodd"/><path d="M13 21.5c-2.834 0-4.5-.208-4.5-1s1.666-1 4.5-1s4.5.208 4.5 1s-1.666 1-4.5 1Z"/><path d="M12.036 16.5h2V21h-2v-4.5Z"/><path d="M17 11a1 1 0 1 1 2 0v1.8c0 2.914-2.721 5.2-6 5.2s-6-2.286-6-5.2V11a1 1 0 0 1 2 0v1.8c0 1.725 1.756 3.2 4 3.2c2.244 0 4-1.475 4-3.2V11Z"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
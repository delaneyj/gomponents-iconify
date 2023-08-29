package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMicrophoneCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M10.5 7.5a2.5 2.5 0 0 1 5 0v4a2.5 2.5 0 0 1-5 0v-4Z" clip-rule="evenodd"/><path d="M13 21.5c-2.834 0-4.5-.208-4.5-1s1.666-1 4.5-1s4.5.208 4.5 1s-1.666 1-4.5 1Z"/><path d="M12.036 16.5h2V21h-2v-4.5Z"/><path d="M17 11a1 1 0 1 1 2 0v1.8c0 2.914-2.721 5.2-6 5.2s-6-2.286-6-5.2V11a1 1 0 0 1 2 0v1.8c0 1.725 1.756 3.2 4 3.2c2.244 0 4-1.475 4-3.2V11Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMicrophoneCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
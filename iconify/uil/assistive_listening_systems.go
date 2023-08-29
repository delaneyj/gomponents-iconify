package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssistiveListeningSystems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a5.07 5.07 0 0 1-.71 2.57L11.73 20A2 2 0 0 1 10 21a2 2 0 0 1-2-2a1 1 0 0 0-2 0a4 4 0 0 0 4 4a4 4 0 0 0 3.5-2.07l3.56-7.43A6.93 6.93 0 0 0 18 10a7 7 0 0 0-7-7ZM4 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM17.59 1.2a1 1 0 1 0-1.2 1.6A9 9 0 0 1 20 10a1 1 0 0 0 2 0a11.06 11.06 0 0 0-4.41-8.8ZM11 9a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-6 0a3 3 0 0 0 .51 1.68a3.5 3.5 0 0 0 .47.54l.2.22a1 1 0 0 1 0 1.11a1 1 0 0 0 .25 1.39a1 1 0 0 0 .57.18a1 1 0 0 0 .82-.43a3 3 0 0 0 0-3.39a3.39 3.39 0 0 0-.35-.42l-.14-.14a1.37 1.37 0 0 1-.16-.18A1 1 0 0 1 10 10a1 1 0 0 1 1-1Zm-4 6a1 1 0 1 0-1-1a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}
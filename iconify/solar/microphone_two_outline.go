package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.25 8a5.75 5.75 0 1 1 11.5 0v3a5.75 5.75 0 0 1-11.5 0V8Zm1.566-.75H9a.75.75 0 0 1 0 1.5H7.75v1.5H9a.75.75 0 0 1 0 1.5H7.816a4.251 4.251 0 0 0 8.368 0H13.5a.75.75 0 0 1 0-1.5h2.75v-1.5H13.5a.75.75 0 0 1 0-1.5h2.684a4.251 4.251 0 0 0-8.368 0ZM4 9.25a.75.75 0 0 1 .75.75v1a7.25 7.25 0 1 0 14.5 0v-1a.75.75 0 0 1 1.5 0v1a8.75 8.75 0 0 1-8 8.718V22a.75.75 0 0 1-1.5 0v-2.282a8.75 8.75 0 0 1-8-8.718v-1A.75.75 0 0 1 4 9.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
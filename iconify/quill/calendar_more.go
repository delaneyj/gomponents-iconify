package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<mask id="quillCalendarMore0" fill="#fff"><path fill-rule="evenodd" d="M12.5 19.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM21 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/></mask><g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2Z"/><path fill="currentColor" fill-rule="evenodd" d="M12.5 19.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM21 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/><path fill="currentColor" d="M11 22a2.5 2.5 0 0 0 2.5-2.5h-2a.5.5 0 0 1-.5.5v2Zm-2.5-2.5A2.5 2.5 0 0 0 11 22v-2a.5.5 0 0 1-.5-.5h-2ZM11 17a2.5 2.5 0 0 0-2.5 2.5h2a.5.5 0 0 1 .5-.5v-2Zm2.5 2.5A2.5 2.5 0 0 0 11 17v2a.5.5 0 0 1 .5.5h2ZM16 22a2.5 2.5 0 0 0 2.5-2.5h-2a.5.5 0 0 1-.5.5v2Zm-2.5-2.5A2.5 2.5 0 0 0 16 22v-2a.5.5 0 0 1-.5-.5h-2ZM16 17a2.5 2.5 0 0 0-2.5 2.5h2a.5.5 0 0 1 .5-.5v-2Zm2.5 2.5A2.5 2.5 0 0 0 16 17v2a.5.5 0 0 1 .5.5h2Zm3 0a.5.5 0 0 1-.5.5v2a2.5 2.5 0 0 0 2.5-2.5h-2ZM21 19a.5.5 0 0 1 .5.5h2A2.5 2.5 0 0 0 21 17v2Zm-.5.5a.5.5 0 0 1 .5-.5v-2a2.5 2.5 0 0 0-2.5 2.5h2Zm.5.5a.5.5 0 0 1-.5-.5h-2A2.5 2.5 0 0 0 21 22v-2Z" mask="url(#quillCalendarMore0)"/></g>`),
		g.Group(children),
	)
}
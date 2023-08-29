package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopBriefcaseCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M6.5 14v4a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-4h2v4a3 3 0 0 1-3 3h-11a3 3 0 0 1-3-3v-4h2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4 10a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-3Zm16 0H6v3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 7.746V9.5H8V7.746a3 3 0 0 1 .09-.727l.061-.247a3 3 0 0 1 2.91-2.272h3.877a3 3 0 0 1 2.91 2.272l.062.247a3 3 0 0 1 .09.727V9.5h-2V7.746a.997.997 0 0 0-.03-.242l-.061-.247a1 1 0 0 0-.97-.757h-3.877a1 1 0 0 0-.97.757l-.062.247a1 1 0 0 0-.03.242Z" clip-rule="evenodd"/><path d="M10.866 14.5a1 1 0 0 1-1.732 0L8.268 13a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M10 13.75a1 1 0 0 1 1 1v1.75a1 1 0 1 1-2 0v-1.75a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path d="M16.866 14.5a1 1 0 0 1-1.732 0l-.866-1.5a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M16 13.75a1 1 0 0 1 1 1v1.75a1 1 0 1 1-2 0v-1.75a1 1 0 0 1 1-1Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopBriefcaseCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
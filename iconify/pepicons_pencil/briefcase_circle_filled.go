package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilBriefcaseCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M6 14.5V18a1.5 1.5 0 0 0 1.5 1.5h11A1.5 1.5 0 0 0 20 18v-3.5h1V18a2.5 2.5 0 0 1-2.5 2.5h-11A2.5 2.5 0 0 1 5 18v-3.5h1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.5 10A1.5 1.5 0 0 1 6 8.5h14a1.5 1.5 0 0 1 1.5 1.5v3a2.5 2.5 0 0 1-2.5 2.5H7A2.5 2.5 0 0 1 4.5 13v-3ZM6 9.5a.5.5 0 0 0-.5.5v3A1.5 1.5 0 0 0 7 14.5h12a1.5 1.5 0 0 0 1.5-1.5v-3a.5.5 0 0 0-.5-.5H6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.5 7.746V9.5h-1V7.746a2.5 2.5 0 0 1 .075-.606l.061-.246A2.5 2.5 0 0 1 11.062 5h3.876a2.5 2.5 0 0 1 2.426 1.894l.061.246c.05.198.075.402.075.606V9.5h-1V7.746c0-.122-.015-.245-.045-.364l-.061-.246A1.5 1.5 0 0 0 14.938 6h-3.876a1.5 1.5 0 0 0-1.456 1.136l-.061.246a1.5 1.5 0 0 0-.045.364Z" clip-rule="evenodd"/><path d="M10.866 14.5a1 1 0 0 1-1.732 0L8.268 13a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M10 14a.75.75 0 0 1 .75.75v1.75a.75.75 0 0 1-1.5 0v-1.75A.75.75 0 0 1 10 14Z" clip-rule="evenodd"/><path d="M16.866 14.5a1 1 0 0 1-1.732 0l-.866-1.5a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M16 14a.75.75 0 0 1 .75.75v1.75a.75.75 0 0 1-1.5 0v-1.75A.75.75 0 0 1 16 14Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilBriefcaseCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
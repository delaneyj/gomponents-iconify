package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 11.5V15a1.5 1.5 0 0 0 1.5 1.5h11A1.5 1.5 0 0 0 17 15v-3.5h1V15a2.5 2.5 0 0 1-2.5 2.5h-11A2.5 2.5 0 0 1 2 15v-3.5h1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M1.5 7A1.5 1.5 0 0 1 3 5.5h14A1.5 1.5 0 0 1 18.5 7v3a2.5 2.5 0 0 1-2.5 2.5H4A2.5 2.5 0 0 1 1.5 10V7ZM3 6.5a.5.5 0 0 0-.5.5v3A1.5 1.5 0 0 0 4 11.5h12a1.5 1.5 0 0 0 1.5-1.5V7a.5.5 0 0 0-.5-.5H3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.5 4.746V6.5h-1V4.746a2.5 2.5 0 0 1 .075-.606l.061-.246A2.5 2.5 0 0 1 8.062 2h3.876a2.5 2.5 0 0 1 2.426 1.894l.061.246c.05.198.075.402.075.606V6.5h-1V4.746c0-.122-.015-.245-.045-.364l-.061-.246A1.5 1.5 0 0 0 11.938 3H8.062a1.5 1.5 0 0 0-1.456 1.136l-.061.246a1.5 1.5 0 0 0-.045.364Z" clip-rule="evenodd"/><path d="M7.866 11.5a1 1 0 0 1-1.732 0L5.268 10a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M7 11a.75.75 0 0 1 .75.75v1.75a.75.75 0 0 1-1.5 0v-1.75A.75.75 0 0 1 7 11Z" clip-rule="evenodd"/><path d="M13.866 11.5a1 1 0 0 1-1.732 0l-.866-1.5a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M13 11a.75.75 0 0 1 .75.75v1.75a.75.75 0 0 1-1.5 0v-1.75A.75.75 0 0 1 13 11Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
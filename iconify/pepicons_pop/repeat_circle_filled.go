package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopRepeatCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M17.519 5.749a.75.75 0 0 1 1.052.13l1.547 1.982a.75.75 0 1 1-1.183.923L17.39 6.8a.75.75 0 0 1 .13-1.052Z"/><path d="M19.983 7.727a.75.75 0 0 0-1.052.14l-1.546 2.017a.75.75 0 1 0 1.19.912l1.547-2.017a.75.75 0 0 0-.14-1.052Z"/><path d="M5.48 12.323a5 5 0 0 1 5-5h7.86a1 1 0 1 1 0 2h-7.86a3 3 0 0 0-3 3v1a1 1 0 1 1-2 0v-1Zm3.008 7.928a.75.75 0 0 1-1.053-.13L5.89 18.14a.75.75 0 1 1 1.182-.923L8.619 19.2a.75.75 0 0 1-.13 1.052Z"/><path d="M6.024 18.273a.75.75 0 0 0 1.051-.14l1.547-2.017a.75.75 0 0 0-1.19-.912L5.884 17.22a.75.75 0 0 0 .139 1.052Z"/><path d="M20.526 13.677a5 5 0 0 1-5 5h-7.86a1 1 0 1 1 0-2h7.86a3 3 0 0 0 3-3v-1a1 1 0 1 1 2 0v1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopRepeatCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M17.519 5.749a.75.75 0 0 1 1.052.13l1.547 1.982a.75.75 0 1 1-1.183.923L17.39 6.8a.75.75 0 0 1 .13-1.052Z"/><path d="M19.983 7.727a.75.75 0 0 0-1.052.14l-1.546 2.017a.75.75 0 1 0 1.19.912l1.547-2.017a.75.75 0 0 0-.14-1.052Z"/><path d="M5.48 12.323a5 5 0 0 1 5-5h7.86a1 1 0 1 1 0 2h-7.86a3 3 0 0 0-3 3v1a1 1 0 1 1-2 0v-1Zm3.008 7.928a.75.75 0 0 1-1.053-.13L5.89 18.14a.75.75 0 1 1 1.182-.923L8.619 19.2a.75.75 0 0 1-.13 1.052Z"/><path d="M6.024 18.273a.75.75 0 0 0 1.051-.14l1.547-2.017a.75.75 0 0 0-1.19-.912L5.884 17.22a.75.75 0 0 0 .139 1.052Z"/><path d="M20.526 13.677a5 5 0 0 1-5 5h-7.86a1 1 0 1 1 0-2h7.86a3 3 0 0 0 3-3v-1a1 1 0 1 1 2 0v1Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}
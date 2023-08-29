package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M15.514 3.749a.75.75 0 0 1 1.052.13l1.547 1.982a.75.75 0 1 1-1.183.923L15.384 4.8a.75.75 0 0 1 .13-1.052Z"/><path d="M17.978 5.727a.75.75 0 0 0-1.052.14L15.38 7.883a.75.75 0 1 0 1.19.912l1.547-2.017a.75.75 0 0 0-.14-1.052Z"/><path d="M3.475 10.323a5 5 0 0 1 5-5h7.86a1 1 0 1 1 0 2h-7.86a3 3 0 0 0-3 3v.5a1 1 0 1 1-2 0v-.5Zm3.008 7.928a.75.75 0 0 1-1.053-.13L3.884 16.14a.75.75 0 1 1 1.182-.923L6.613 17.2a.75.75 0 0 1-.13 1.052Z"/><path d="M4.019 16.273a.75.75 0 0 0 1.051-.14l1.547-2.017a.75.75 0 0 0-1.19-.912L3.88 15.22a.75.75 0 0 0 .139 1.052Z"/><path d="M18.521 11.677a5 5 0 0 1-5 5h-7.86a1 1 0 1 1 0-2h7.86a3 3 0 0 0 3-3v-.5a1 1 0 1 1 2 0v.5Z"/></g><path fill-rule="evenodd" d="M14.667 2.946a.5.5 0 0 1 .702.087l1.547 1.982a.5.5 0 0 1-.789.615l-1.546-1.982a.5.5 0 0 1 .086-.702Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.826 4.926a.5.5 0 0 0-.701.092l-1.547 2.018a.5.5 0 0 0 .794.608l1.546-2.017a.5.5 0 0 0-.092-.701Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.975 9.323a4.5 4.5 0 0 1 4.5-4.5h8.5a.5.5 0 0 1 0 1h-8.5a3.5 3.5 0 0 0-3.5 3.5v1.337a.5.5 0 1 1-1 0V9.323Zm2.354 7.731a.5.5 0 0 1-.702-.087l-1.546-1.982a.5.5 0 0 1 .788-.615l1.547 1.982a.5.5 0 0 1-.087.702Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.17 15.074a.5.5 0 0 0 .702-.092l1.546-2.018a.5.5 0 1 0-.793-.608l-1.547 2.017a.5.5 0 0 0 .093.701Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.021 10.677a4.5 4.5 0 0 1-4.5 4.5h-8.5a.5.5 0 0 1 0-1h8.5a3.5 3.5 0 0 0 3.5-3.5V9.34a.5.5 0 1 1 1 0v1.337Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
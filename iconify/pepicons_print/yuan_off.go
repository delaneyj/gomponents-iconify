package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YuanOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M4.919 2.936a1 1 0 0 1 1.395.233l5 7a1 1 0 1 1-1.628 1.162l-5-7a1 1 0 0 1 .233-1.395Z"/><path d="M16.081 2.936a1 1 0 0 0-1.395.233l-5 7a1 1 0 0 0 1.628 1.162l5-7a1 1 0 0 0-.233-1.395Z"/><path d="M5 11.75a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1Z"/><path d="M10.5 9.75a1 1 0 0 1 1 1v7a1 1 0 1 1-2 0v-7a1 1 0 0 1 1-1Z"/></g><path fill-rule="evenodd" d="M3.832 2.584a.5.5 0 0 1 .694.139l5 7.5a.5.5 0 0 1-.832.554l-5-7.5a.5.5 0 0 1 .138-.693Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.387 2.584a.5.5 0 0 0-.693.139l-5 7.5a.5.5 0 0 0 .832.554l5-7.5a.5.5 0 0 0-.139-.693Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.11 11a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.11 10a.5.5 0 0 1 .5.5V17a.5.5 0 0 1-1 0v-6.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
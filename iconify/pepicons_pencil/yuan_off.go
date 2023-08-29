package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YuanOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.832 2.584a.5.5 0 0 1 .694.139l5 7.5a.5.5 0 0 1-.832.554l-5-7.5a.5.5 0 0 1 .138-.693Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.387 2.584a.5.5 0 0 0-.693.139l-5 7.5a.5.5 0 0 0 .832.554l5-7.5a.5.5 0 0 0-.139-.693Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.11 11a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.11 10a.5.5 0 0 1 .5.5V17a.5.5 0 0 1-1 0v-6.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
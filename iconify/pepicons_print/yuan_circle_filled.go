package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YuanCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M6.832 5.584a.5.5 0 0 1 .694.139l5 7.5a.5.5 0 0 1-.832.554l-5-7.5a.5.5 0 0 1 .138-.693Z"/><path d="M17.387 5.584a.5.5 0 0 0-.693.139l-5 7.5a.5.5 0 0 0 .832.554l5-7.5a.5.5 0 0 0-.139-.693Z"/><path d="M7.11 14a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Z"/><path d="M12.11 13a.5.5 0 0 1 .5.5V20a.5.5 0 0 1-1 0v-6.5a.5.5 0 0 1 .5-.5Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
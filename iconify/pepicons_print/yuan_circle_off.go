package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YuanCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M7.919 5.936a1 1 0 0 1 1.395.233l5 7a1 1 0 1 1-1.628 1.162l-5-7a1 1 0 0 1 .233-1.395Z"/><path d="M19.081 5.936a1 1 0 0 0-1.395.233l-5 7a1 1 0 0 0 1.628 1.162l5-7a1 1 0 0 0-.233-1.395Z"/><path d="M8 14.75a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1Z"/><path d="M13.5 12.75a1 1 0 0 1 1 1v7a1 1 0 1 1-2 0v-7a1 1 0 0 1 1-1Z"/></g><path fill-rule="evenodd" d="M6.832 5.584a.5.5 0 0 1 .694.139l5 7.5a.5.5 0 0 1-.832.554l-5-7.5a.5.5 0 0 1 .138-.693Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.387 5.584a.5.5 0 0 0-.693.139l-5 7.5a.5.5 0 0 0 .832.554l5-7.5a.5.5 0 0 0-.139-.693Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.11 14a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.11 13a.5.5 0 0 1 .5.5V20a.5.5 0 0 1-1 0v-6.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
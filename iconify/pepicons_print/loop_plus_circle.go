package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopPlusCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M9.828 8.828a6 6 0 0 0 6.865 9.642c.058.091.127.178.206.258l3.536 3.535a1.5 1.5 0 0 0 2.121-2.12l-3.535-3.536l-.06-.057a6.002 6.002 0 0 0-9.133-7.722Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M8.182 8.182a5.5 5.5 0 1 0 7.778 7.778a5.5 5.5 0 0 0-7.778-7.778Zm7.071 7.071A4.5 4.5 0 1 1 8.89 8.889a4.5 4.5 0 0 1 6.364 6.364Z" clip-rule="evenodd"/><path d="M15 17.121a1 1 0 0 1 1.414-1.414l3.789 3.789a1 1 0 0 1-1.414 1.414L15 17.12Zm-4.793-4.457a.5.5 0 1 1 0-1h4a.5.5 0 1 1 0 1h-4Z"/><path d="M11.707 10.164a.5.5 0 0 1 1 0v4a.5.5 0 0 1-1 0v-4Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
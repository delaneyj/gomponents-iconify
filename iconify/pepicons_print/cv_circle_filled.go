package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CvCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill="currentColor" d="M9.5 15.5a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1h-7Zm0 2.5a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1h-7Z"/><path fill="currentColor" fill-rule="evenodd" d="M14.185 4H7.5A1.5 1.5 0 0 0 6 5.5v15A1.5 1.5 0 0 0 7.5 22h11a1.5 1.5 0 0 0 1.5-1.5V10.202a1.5 1.5 0 0 0-.395-1.014l-4.314-4.702A1.5 1.5 0 0 0 14.185 4ZM7 5.5a.5.5 0 0 1 .5-.5h6.685a.5.5 0 0 1 .369.162l4.314 4.702a.5.5 0 0 1 .132.338V20.5a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-15Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 5.1v4.7h4.7"/><path fill="currentColor" d="M11.134 9.133a1.067 1.067 0 1 0 0-2.133a1.067 1.067 0 0 0 0 2.133Z"/><path fill="currentColor" fill-rule="evenodd" d="M13.266 11.444c0-1.134-.955-1.955-2.133-1.955S9 10.309 9 11.444v.534a.356.356 0 0 0 .356.355h3.555a.356.356 0 0 0 .355-.355v-.534Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M17.052 10.852a.5.5 0 0 1 .453.543l-.472 5.186a.5.5 0 1 1-.995-.09l.47-5.186a.5.5 0 0 1 .544-.453Z"/><path d="M10.852 17.052a.5.5 0 0 1 .453-.543l5.185-.471a.5.5 0 0 1 .09.995l-5.185.472a.5.5 0 0 1-.543-.453Z"/><path d="M16.354 16.354a.5.5 0 0 1-.708 0L9.111 9.818a.5.5 0 1 1 .707-.707l6.536 6.535a.5.5 0 0 1 0 .708Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
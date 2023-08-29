package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.5 14v4a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-4h2v4a3 3 0 0 1-3 3h-11a3 3 0 0 1-3-3v-4h2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4 10a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-3Zm16 0H6v3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 7.746V9.5H8V7.746a3 3 0 0 1 .09-.727l.061-.247a3 3 0 0 1 2.91-2.272h3.877a3 3 0 0 1 2.91 2.272l.062.247a3 3 0 0 1 .09.727V9.5h-2V7.746a.997.997 0 0 0-.03-.242l-.061-.247a1 1 0 0 0-.97-.757h-3.877a1 1 0 0 0-.97.757l-.062.247a1 1 0 0 0-.03.242Z" clip-rule="evenodd"/><path d="M10.866 14.5a1 1 0 0 1-1.732 0L8.268 13a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M10 13.75a1 1 0 0 1 1 1v1.75a1 1 0 1 1-2 0v-1.75a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path d="M16.866 14.5a1 1 0 0 1-1.732 0l-.866-1.5a1 1 0 0 1 .866-1.5h1.732a1 1 0 0 1 .866 1.5l-.866 1.5Z"/><path fill-rule="evenodd" d="M16 13.75a1 1 0 0 1 1 1v1.75a1 1 0 1 1-2 0v-1.75a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteSingleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M9.781 1.698c4.323 2.937 6.469 5.064 6.469 6.606c0 1.492-.82 2.7-2.396 3.583c-.436.245-.922-.232-.685-.672c.407-.758.273-1.607-.461-2.617c-.774-1.065-1.89-1.84-3.365-2.328A.5.5 0 0 1 9 5.795V2.111a.5.5 0 0 1 .781-.413ZM7.75 17.75c-1.77 0-3.25-1.143-3.25-2.625S5.98 12.5 7.75 12.5S11 13.643 11 15.125S9.52 17.75 7.75 17.75Z"/><path d="M10 4a1 1 0 0 1 1 1v10a1 1 0 1 1-2 0V5a1 1 0 0 1 1-1Z"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}
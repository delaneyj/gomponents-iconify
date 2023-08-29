package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleUpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path d="m19.32 15.116l-6-5l-.64.768l6 5l.64-.768Z"/><path d="m18.68 15.884l-6-5c-.512-.427.128-1.195.64-.768l6 5c.512.427-.128 1.195-.64.768Z"/><path d="m7.32 15.884l6-5l-.64-.768l-6 5l.64.768Z"/><path d="m6.68 15.116l6-5c.512-.427 1.152.341.64.768l-6 5c-.512.427-1.152-.341-.64-.768Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
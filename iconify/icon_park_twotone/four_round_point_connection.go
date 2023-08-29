package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourRoundPointConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFourRoundPointConnection0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M10 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M6.283 14.71A19.859 19.859 0 0 0 4.4 20M14.71 6.283A19.86 19.86 0 0 1 20 4.4"/><path fill="#555" d="M38 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M33.29 6.283A19.86 19.86 0 0 0 28 4.4m13.716 10.31A19.857 19.857 0 0 1 43.6 20"/><path fill="#555" d="M38 44a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M41.716 33.29A19.859 19.859 0 0 0 43.6 28M33.29 41.716A19.859 19.859 0 0 1 28 43.6"/><path fill="#555" d="M10 44a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M14.71 41.716A19.859 19.859 0 0 0 20 43.6M6.283 33.29A19.86 19.86 0 0 1 4.4 28M24 30V18m-6 6h12h-12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFourRoundPointConnection0)"/>`),
		g.Group(children),
	)
}
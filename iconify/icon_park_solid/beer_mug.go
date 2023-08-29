package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerMug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBeerMug0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M16 37.241C16 39.172 16.857 44 28 44s12-4.828 12-6.759V16H16v21.241Z"/><path stroke="#000" stroke-linecap="round" d="M25 23v14m6-14v14"/><path stroke="#fff" d="M15.998 16s-1.999-4.5 1-7C19.999 6.5 23 8 23 8s1-4 5-4s5 4 5 4s3.5-1.5 6 1s.998 7 .998 7M16 19H6s1 10 2 13c.998 3 8 2 8 2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBeerMug0)"/>`),
		g.Group(children),
	)
}
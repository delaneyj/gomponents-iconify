package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H3v6h18V3m-2 4H5V5h14v2m-4.5 4c.28 0 .5.22.5.5V13H9v-1.5c0-.28.22-.5.5-.5h5m3.5 2.09V10h2v3.09c-.33-.05-.66-.09-1-.09c-.34 0-.67.04-1 .09M13 19c0 .7.13 1.37.35 2H4V10h2v9h7m9.5-1.75L17.75 22L15 19l1.16-1.16l1.59 1.59l3.59-3.59l1.16 1.41Z"/>`),
		g.Group(children),
	)
}
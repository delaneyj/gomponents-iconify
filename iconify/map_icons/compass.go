package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25 49C11.766 49 1 38.233 1 25C1 11.766 11.766 1 25 1c13.233 0 24 10.766 24 24c0 13.233-10.767 24-24 24zm0-44C13.972 5 5 13.972 5 25s8.972 20 20 20s20-8.972 20-20S36.028 5 25 5zm.045 3.25S18 20.321 18 24v2c0 3.678 7.066 16 7.066 16S32 29.934 32 26.256v-2.262c0-3.679-6.955-15.744-6.955-15.744zM25 29a4 4 0 1 1 0-8a4 4 0 0 1 0 8z"/>`),
		g.Group(children),
	)
}
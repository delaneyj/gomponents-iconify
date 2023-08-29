package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMaximum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm1.998.998h2.004v2.004H5.998V2.998Zm3 0h2.004v2.004H8.998V2.998Zm3 0h2.004v2.004h-2.004V2.998Zm3 0h2.004v2.004h-2.004V2.998Zm3 0h2.004v2.004h-2.004V2.998Zm-2.567 5.789C14.816 8.304 14.035 7.999 13 8c-1.034 0-1.816.306-2.432.79c-.63.495-1.136 1.218-1.53 2.116C8.247 12.721 8 15.059 8 16.998v1H6v-1c0-2.047.252-4.707 1.207-6.893c.482-1.102 1.163-2.131 2.127-2.888C10.312 6.448 11.53 6 12.999 6c1.468-.001 2.688.445 3.667 1.214c.964.757 1.645 1.787 2.127 2.89c.955 2.189 1.207 4.852 1.207 6.9v1h-2v-1c0-1.94-.246-4.282-1.04-6.1c-.393-.899-.899-1.622-1.53-2.117Z"/>`),
		g.Group(children),
	)
}
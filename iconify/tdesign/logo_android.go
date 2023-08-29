package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.372 3.636l2.116 3.702A11.952 11.952 0 0 1 12 6c1.986 0 3.86.483 5.512 1.338l2.116-3.702l1.736.992l-2.158 3.776A11.983 11.983 0 0 1 24 18v2H0v-2c0-3.924 1.884-7.407 4.793-9.596L2.636 4.628l1.736-.992ZM12 8a9.948 9.948 0 0 0-5.348 1.55A9.992 9.992 0 0 0 2 18h20a9.992 9.992 0 0 0-4.652-8.451A9.949 9.949 0 0 0 12 8Zm-6 5h2.004v2.004H6V13Zm10 0h2.004v2.004H16V13Z"/>`),
		g.Group(children),
	)
}
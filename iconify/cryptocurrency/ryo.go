package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ryo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-.014-28C9.376 4 4 9.376 4 15.986S9.376 27.97 15.986 27.97S27.97 22.595 27.97 15.986C27.971 9.376 22.595 4 15.986 4zm0 1.632a10.34 10.34 0 0 1 10.353 10.354a10.34 10.34 0 0 1-10.353 10.353A10.34 10.34 0 0 1 5.632 15.986A10.341 10.341 0 0 1 15.986 5.632zm-4.453 5.928v8.85h8.905v-8.85h-8.905zm1.45 1.45h6.004v5.95h-6.003v-5.95z"/>`),
		g.Group(children),
	)
}
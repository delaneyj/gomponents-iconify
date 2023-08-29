package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brightness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBrightness0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24.02 35.353c6.076 0 11-4.924 11-11c0-6.075-4.924-11-11-11c-6.075 0-11 4.925-11 11c0 6.076 4.925 11 11 11Z"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="m38.96 9.01l-2.46 2.474M11.067 36.745l-2.046 2.059M24 41.353v3m20-20.999h-4m-2.468 12.982L40 38.804"/><path fill="#fff" fill-rule="evenodd" d="M24.02 17.354a7 7 0 1 0 0 14" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M4 24.354h4M10.044 9.01l2.053 2.053M24 3.354v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBrightness0)"/>`),
		g.Group(children),
	)
}
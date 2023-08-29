package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PyramidOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPyramidOne0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" fill-rule="evenodd" stroke-linejoin="round" d="m19 14l14 28H5l14-28Z" clip-rule="evenodd"/><path stroke-linecap="round" stroke-linejoin="round" d="m24 25l-9 17m18 0h10.5L35 28l-4.5 8"/><path stroke-linecap="round" d="M25.984 26.396c6.228-1.582 9.994-7.914 8.412-14.142c-1.582-6.228-7.914-9.995-14.142-8.412c-6.229 1.582-9.995 7.913-8.413 14.142a11.604 11.604 0 0 0 1.937 4.078"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 42H10m12.5-21l5 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPyramidOne0)"/>`),
		g.Group(children),
	)
}
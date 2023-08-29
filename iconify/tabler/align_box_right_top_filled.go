package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBoxRightTopFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18.333 2c1.96 0 3.56 1.537 3.662 3.472l.005.195v12.666c0 1.96-1.537 3.56-3.472 3.662l-.195.005H5.667a3.667 3.667 0 0 1-3.662-3.472L2 18.333V5.667c0-1.96 1.537-3.56 3.472-3.662L5.667 2h12.666zM18 11h-2l-.117.007A1 1 0 0 0 16 13h2l.117-.007A1 1 0 0 0 18 11zm0-3h-6l-.117.007A1 1 0 0 0 12 10h6l.117-.007A1 1 0 0 0 18 8zm0-3h-4l-.117.007A1 1 0 0 0 14 7h4l.117-.007A1 1 0 0 0 18 5z"/></g>`),
		g.Group(children),
	)
}
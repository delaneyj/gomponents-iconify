package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMountain0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m18 10l14 28H4l14-28Z"/><path stroke-linecap="round" d="m28 29l5.647-7L44 38H32M12 22h12m-10-4l-4 8m12-8l4 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMountain0)"/>`),
		g.Group(children),
	)
}
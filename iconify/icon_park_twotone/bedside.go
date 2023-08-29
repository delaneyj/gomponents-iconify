package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bedside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBedside0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 18h40v12H4zm0 12h40v12H4z"/><path d="M22 24h4m-4 12h4M8 42v4m32-4v4M24 18v-8m8 0H16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBedside0)"/>`),
		g.Group(children),
	)
}
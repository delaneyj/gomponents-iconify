package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentBanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTentBanner0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M7 43h32L23 17L7 43Z"/><path d="M18.077 25L23 28l4.923-3M23 17V4"/><path fill="#555" d="M35 4H23v8h12l-3-4l3-4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTentBanner0)"/>`),
		g.Group(children),
	)
}
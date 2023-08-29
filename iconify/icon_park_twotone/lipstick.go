package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLipstick0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M29 24h12v20H29zM7 24h14v20H7z"/><path fill="#555" d="M10 11.454V24h8V4c-6.5 0-8 5.636-8 7.454Z"/><path d="M7 32h14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLipstick0)"/>`),
		g.Group(children),
	)
}
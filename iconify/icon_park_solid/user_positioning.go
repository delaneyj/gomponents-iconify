package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPositioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUserPositioning0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="16" r="6" fill="#fff"/><path d="M36 36c0-6.627-5.373-12-12-12s-12 5.373-12 12M36 4h8v8M12 4H4v8m32 32h8v-8m-32 8H4v-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUserPositioning0)"/>`),
		g.Group(children),
	)
}
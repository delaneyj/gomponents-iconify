package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSettingOne0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="m34 41l10-17L34 7H14L4 24l10 17h20Z"/><path d="M24 29a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSettingOne0)"/>`),
		g.Group(children),
	)
}
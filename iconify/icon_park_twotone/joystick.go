package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joystick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTJoystick0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 33H4v10h40V33Z"/><path stroke-linecap="round" d="M38 26H26v7h12v-7Z"/><path fill="#555" d="M18 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" d="M16 14L9 33"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTJoystick0)"/>`),
		g.Group(children),
	)
}
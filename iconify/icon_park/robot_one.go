package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="30" height="24" x="9" y="18" fill="#2F88FF" stroke="#000" stroke-width="4" rx="2"/><circle cx="17" cy="26" r="2" fill="#fff"/><circle cx="31" cy="26" r="2" fill="#fff"/><path fill="#fff" d="M20 32C18.8954 32 18 32.8954 18 34C18 35.1046 18.8954 36 20 36V32ZM28 36C29.1046 36 30 35.1046 30 34C30 32.8954 29.1046 32 28 32V36ZM20 36H28V32H20V36Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 10V18"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 26V34"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 26V34"/><circle cx="24" cy="8" r="2" stroke="#000" stroke-width="4"/></g>`),
		g.Group(children),
	)
}
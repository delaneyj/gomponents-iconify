package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handwashing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="7" height="20" x="4" y="24"/><path fill="#2F88FF" d="M32.0003 16C30.5006 12.5 36.0003 6 36.0003 6C36.0003 6 41.5001 12.5 40.0003 16C38.5005 19.5 33.5 19.5 32.0003 16Z"/><path fill="#2F88FF" d="M31 42.5001C22 42.5001 15 40.0001 11 40.0001V30.0001C18 30.0001 17.5 27.5001 22 26.0001C26.5 24.5001 30 26.0001 29.5 29.0001C29 32.0001 24 35.0001 24 35.0001C32 35.0001 32 33.0001 36 30.0001C40 27.0001 44 28.0001 44 31.0001C44 34.0001 40 42.5001 31 42.5001Z"/></g>`),
		g.Group(children),
	)
}
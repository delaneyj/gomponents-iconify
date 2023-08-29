package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaskTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="24" r="20"/><path stroke-linecap="round" stroke-linejoin="round" d="M31 5.25928C23.4067 8.09675 18 15.417 18 24.0001C18 32.5831 23.4067 39.9034 31 42.7408"/><path stroke-linecap="round" stroke-linejoin="round" d="M37 9L18 22"/><path stroke-linecap="round" stroke-linejoin="round" d="M41 14L19 29"/><path stroke-linecap="round" stroke-linejoin="round" d="M43 20L22 35"/><path stroke-linecap="round" stroke-linejoin="round" d="M43 28L26 40"/></g>`),
		g.Group(children),
	)
}
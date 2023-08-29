package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsWheelchairAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm16 42c5.523 0 10-4.477 10-10a9.966 9.966 0 0 0-.832-4h5.59l.242.242v7.929a3.001 3.001 0 1 0 2 0v-5.929l5.364 5.365l.707.707l.707-.707l2.829-2.829l-1.414-1.414l-2.122 2.121L32 26.414V26a2 2 0 0 0-2-2h-1v-4a3 3 0 0 0-3-3H15.53L12.97 6.757L12.78 6H8v2h3.22l3.519 14.079C9.81 22.699 6 26.904 6 32c0 5.523 4.477 10 10 10Zm.03-23H26a1 1 0 0 1 1 1v4h-4.999a9.95 9.95 0 0 0-5.213-1.97L16.031 19ZM16 37.556a5.556 5.556 0 1 0 0-11.112a5.556 5.556 0 0 0 0 11.112ZM31 39a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm-14-7a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsWheelchairAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
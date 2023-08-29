package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m657 679l-272 345l-272-345q-73-76-99.5-178.5t0-205T113 116q56-57 127.5-86.5T385 0t144.5 29.5T657 116q73 77 99.5 179.5t0 205T657 679zM385.5 129Q279 129 204 204t-75 181t75 180.5T385.5 640T567 565.5T642 385t-75-181t-181.5-75zm.5 383q-53 0-90.5-37.5T258 384t37.5-90.5T386 256t90.5 37.5T514 384t-37.5 90.5T386 512z"/>`),
		g.Group(children),
	)
}
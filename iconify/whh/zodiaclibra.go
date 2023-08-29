package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zodiaclibra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 768H640q-27 0-45.5-19T576 703.5t18.5-45T640 640q15 0 46.5-50T743 477.5t25-93.5q0-106-75-181t-181-75t-181 75t-75 181q0 31 25 93.5T337.5 590t46.5 50q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19H64q-27 0-45.5-19T0 703.5t18.5-45T64 640h151q-39-61-63-131t-24-125q0-104 51.5-192.5t140-140T512 0t192.5 51.5t140 140T896 384q0 55-24 125t-63 131h151q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zM64 896h896q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19H64q-27 0-45.5-19T0 959.5t18.5-45T64 896z"/>`),
		g.Group(children),
	)
}
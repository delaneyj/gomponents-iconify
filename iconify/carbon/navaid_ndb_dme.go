package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidNdbDme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 12H2V4a2.002 2.002 0 0 1 2-2h8v2H4zm8 18H4a2.002 2.002 0 0 1-2-2v-8h2v8h8zm16 0h-8v-2h8v-8h2v8a2.002 2.002 0 0 1-2 2zm2-18h-2V4h-8V2h8a2.002 2.002 0 0 1 2 2zM16 22a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6zm0-10a4 4 0 1 0 4 4a4.005 4.005 0 0 0-4-4z"/><circle cx="16" cy="16" r="2" fill="currentColor"/><circle cx="16" cy="7" r="1" fill="currentColor"/><circle cx="16" cy="3" r="1" fill="currentColor"/><circle cx="16" cy="25" r="1" fill="currentColor"/><circle cx="16" cy="29" r="1" fill="currentColor"/><circle cx="25" cy="16" r="1" fill="currentColor"/><circle cx="29" cy="16" r="1" fill="currentColor"/><circle cx="7" cy="16" r="1" fill="currentColor"/><circle cx="3" cy="16" r="1" fill="currentColor"/><circle cx="9.636" cy="9.636" r="1" fill="currentColor"/><circle cx="7" cy="7" r="1" fill="currentColor"/><circle cx="22.364" cy="22.364" r="1" fill="currentColor"/><circle cx="25" cy="25" r="1" fill="currentColor"/><circle cx="22.364" cy="9.636" r="1" fill="currentColor"/><circle cx="25" cy="7" r="1" fill="currentColor"/><circle cx="9.636" cy="22.364" r="1" fill="currentColor"/><circle cx="7" cy="25" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}
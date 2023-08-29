package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Itsme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.316 31.756a3.611 3.611 0 0 1 7.223 0v5.778m-7.223-9.389v9.389"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.538 31.756a3.611 3.611 0 0 1 7.223 0v5.778m8.985-1.806a3.495 3.495 0 0 1-3.07 1.805a3.622 3.622 0 0 1-3.611-3.61v-2.348a3.611 3.611 0 0 1 7.222 0v1.264h-7.222"/><circle cx="12.405" cy="11.697" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.69 23.238a4.574 4.574 0 0 0 3.095.774h.775a2.516 2.516 0 0 0 0-5.031h-1.742a2.516 2.516 0 0 1 0-5.031h.774c1.742 0 2.516.193 3.096.774"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="8.148" d="M12.4 22.077a1.827 1.827 0 0 0 1.934 1.935a3.595 3.595 0 0 0 2.613-.774m.194-12.772v11.61a1.827 1.827 0 0 0 1.935 1.936a3.595 3.595 0 0 0 2.613-.774"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.141 13.756h2.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="8.148" d="m12.4 22.077l.006-8.32"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
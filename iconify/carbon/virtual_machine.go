package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 13h-8a2.002 2.002 0 0 0-2 2v13a2.002 2.002 0 0 0 2 2h8a2.002 2.002 0 0 0 2-2V15a2.002 2.002 0 0 0-2-2Zm0 2v3h-8v-3Zm0 5v3h-8v-3Zm-8 8v-3h8v3Z"/><path fill="currentColor" d="M25.91 10.13a.121.121 0 0 1-.097-.095a10.006 10.006 0 0 0-7.923-7.853a10.002 10.002 0 0 0-11.704 7.853a.121.121 0 0 1-.097.095A7.502 7.502 0 0 0 7.491 25H8v-2h-.505a5.51 5.51 0 0 1-5.438-6.3a5.7 5.7 0 0 1 4.713-4.66l1.017-.184a.13.13 0 0 0 .105-.104l.18-.935a8.28 8.28 0 0 1 6.847-6.743a7.957 7.957 0 0 1 2.847.125a8.22 8.22 0 0 1 6.147 6.545l.194 1.008a.13.13 0 0 0 .105.104l1.057.191a5.782 5.782 0 0 1 3.101 1.54A5.505 5.505 0 0 1 24.508 23H24v2h.508a7.502 7.502 0 0 0 1.402-14.87Z"/>`),
		g.Group(children),
	)
}
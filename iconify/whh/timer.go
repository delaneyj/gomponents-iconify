package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512q0-120 54-229l115 57q-41 82-41 172q0 104 51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512q0-122-69.5-220T647 153l45-120q147 55 239.5 186t92.5 293q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-448q-36 0-55-32q-32-34-83-93T261.5 319.5t-101-128t-32-63t63 32t128 101t131.5 113t93 82.5q32 19 32 55q0 27-18.5 45.5T512 576zm0-448q-27 0-45.5-19T448 63.5t18.5-45t45-18.5T557 18.5T576 64t-18.5 45.5T512 128z"/>`),
		g.Group(children),
	)
}
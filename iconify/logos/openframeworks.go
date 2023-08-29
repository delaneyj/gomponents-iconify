package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openframeworks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M67.49 0c37.237 0 67.492 30.255 67.492 67.49c0 37.237-30.255 67.492-67.491 67.492C30.255 134.982 0 104.727 0 67.49C0 30.255 30.255 0 67.49 0Zm128 0v134.982h-58.18V0h58.18Zm37.237 60.51v34.908h-34.909V60.51h34.91ZM256 0l-58.182 58.182V0H256Z"/>`),
		g.Group(children),
	)
}
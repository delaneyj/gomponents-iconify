package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceTabletBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 20H64a28 28 0 0 0-28 28v160a28 28 0 0 0 28 28h128a28 28 0 0 0 28-28V48a28 28 0 0 0-28-28ZM60 84h136v88H60Zm4-40h128a4 4 0 0 1 4 4v12H60V48a4 4 0 0 1 4-4Zm128 168H64a4 4 0 0 1-4-4v-12h136v12a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}
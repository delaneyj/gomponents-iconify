package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moonwaninggibbous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-960q61 0 122 72t97.5 176.5T768 512t-36.5 199.5T634 888t-122 72q91 0 174-35.5T829 829t95.5-143T960 512t-35.5-174T829 195T686 99.5T512 64z"/>`),
		g.Group(children),
	)
}
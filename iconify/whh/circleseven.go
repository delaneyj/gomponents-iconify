package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleseven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm64-768H352q-13 0-22.5 9.5T320 288t9.5 22.5T352 320h224q27 0 45.5 19t18.5 45q4 43-16 80q-28 52-128 240q-16 22-16 32q0 12 10 22t22 10q13 0 22.5-9.5T544 736l144-272q16-30 16-80q0-53-37.5-90.5T576 256z"/>`),
		g.Group(children),
	)
}
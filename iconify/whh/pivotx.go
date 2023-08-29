package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pivotx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m960 361l-113 35q69 81 97 167.5t6 140.5l-88-36q5 90-17 160.5T782 928l-69-95q-50 87-114 139t-119 52V914q-87 41-166.5 45.5T192 931l75-110q-102-18-175.5-62T0 662l112-35q-68-81-96-167T9 320l89 36q-5-90 17-160.5T177 96l70 95q50-87 114-139T480 0v110q87-41 166.5-45.5T768 93l-75 109q102 18 175.5 62t91.5 97zM480 256q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75z"/>`),
		g.Group(children),
	)
}
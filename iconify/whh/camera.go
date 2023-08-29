package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 960H128q-53 0-90.5-37.5T0 832V320q0-53 37.5-90.5T128 192h93l42-94q11-42 55-70t98-28h192q54 0 98 28t54 70l43 94h93q53 0 90.5 37.5T1024 320v512q0 53-37.5 90.5T896 960zM512 256q-87 0-160.5 43T235 415.5T192 576t43 160.5T351.5 853T512 896t160.5-43T789 736.5T832 576t-43-160.5T672.5 299T512 256zm0 576q-106 0-181-75t-75-181t75-181t181-75t181 75t75 181t-75 181t-181 75zm0-448q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136t-56-136t-136-56z"/>`),
		g.Group(children),
	)
}
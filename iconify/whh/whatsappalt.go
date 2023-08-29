package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whatsappalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M544 960q-124 0-233-60L0 1024l124-311Q64 604 64 480q0-98 38-186.5t102.5-153T357.5 38T544 0t186.5 38t153 102.5t102.5 153t38 186.5t-38 186.5t-102.5 153t-153 102.5T544 960zm0-832q-96 0-177 47T239 303t-47 177t47 177t128 128t177 47t177-47t128-128t47-177t-47-177t-128-128t-177-47zm173 554q-39 39-121.5 11.5t-160-105t-105-160T342 307l47-47q6-6 20.5 1.5T436 286t12 34v64l-32 36q12 45 77.5 110.5T604 608l36-32h64q17 0 34 12t24.5 26.5T764 635z"/>`),
		g.Group(children),
	)
}
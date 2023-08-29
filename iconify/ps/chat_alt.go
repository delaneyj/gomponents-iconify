package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M463 226q6 20 6 36q0 63-68 107l-17 15q4 64 9 77q-12-6-54-36q-18-20-38-20h-64q-9 9-43 22q53 21 107 21q1 0 19 16t48.5 32t60.5 16q7 0 9-11t0-26.5t-4.5-31.5t-5.5-27l-3-11q89-56 89-143q0-67-43-113q-1 7-.5 20t-1 31.5T463 226zM79 427q30 0 62-21.5t51.5-43T213 341q90 0 152-46t62-118q0-73-62.5-125T213 0Q125 0 62.5 52T0 177q0 86 90 143q-2 7-5.5 18T75 372.5t-5.5 39T79 427zM43 177q0-56 49.5-95T213 43t121 39t50 95t-51 89t-120 33q-22 0-42 25q-35 35-52 45q1-2 11-36l11-32l-28-17q-70-44-70-107z"/>`),
		g.Group(children),
	)
}
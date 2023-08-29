package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yui(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 960q-95 0-197.5-34.5T256 832q12-34 29-71q-104-68-157-121q14-24 34-54Q78 526 0 448q75-125 229.5-227T593 59.5T1024 0Q640 480 640 960zm-64-64q0-173 79-381t209-387q-38 26-109.5 97.5T603.5 391t-151 206.5T352 800q62 48 113.5 72T576 896zM96 448q52 48 103 82q202-297 441-429q-94 27-185 72t-160 94t-119 95.5T96 448zm128 192q32 32 87 66q88-174 231.5-347.5T822 87q-75 23-152 71.5T530 264T406.5 391T303 520.5T224 640z"/>`),
		g.Group(children),
	)
}
package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.15 4a1.85 1.85 0 1 1 3.7 0a1.85 1.85 0 0 1-3.7 0ZM4 1.25a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5ZM5.82 11L2.5 12.837V9.163L5.82 11ZM2.64 8.212a.7.7 0 0 0-1.039.612v4.352a.7.7 0 0 0 1.039.613l3.933-2.176a.7.7 0 0 0 0-1.225L2.64 8.212ZM8.3 9a.7.7 0 0 1 .7-.7h4a.7.7 0 0 1 .7.7v4a.7.7 0 0 1-.7.7H9a.7.7 0 0 1-.7-.7V9Zm.9.2v3.6h3.6V9.2H9.2Zm4.243-7.007a.45.45 0 0 0-.636-.636L11 3.364L9.193 1.557a.45.45 0 1 0-.636.636L10.364 4L8.557 5.807a.45.45 0 1 0 .636.636L11 4.636l1.807 1.807a.45.45 0 0 0 .636-.636L11.636 4l1.807-1.807Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
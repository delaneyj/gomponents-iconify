package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BfFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#de0000" d="M640 479.6H.4V0H640z"/><path fill="#35a100" d="M639.6 480H0V240.2h639.6z"/><path fill="#fff300" d="m254.6 276.2l-106-72.4h131L320 86.6L360.4 204l131-.1l-106 72.4l40.5 117.3l-106-72.6L214 393.4"/></g>`),
		g.Group(children),
	)
}
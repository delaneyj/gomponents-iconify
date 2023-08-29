package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JmFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path d="m0 0l320 240L0 480zm640 0L320 240l320 240z"/><path fill="#090" d="m0 0l320 240L640 0zm0 480l320-240l320 240z"/><path fill="#fc0" d="M640 0h-59.6L0 435.3V480h59.6L640 44.7z"/><path fill="#fc0" d="M0 0v44.7L580.4 480H640v-44.7L59.6 0z"/></g>`),
		g.Group(children),
	)
}
package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 351.3h85.3v-64H0v64zM0 138h85.3V74H0v64zm0 320h85.3v-64H0v64zm0-213.3h85.3v-64H0v64zM170.7 458H512v-64H170.7v64zm0-384v64H512V74H170.7zm0 170.7H512v-64H170.7v64zm0 106.6H512v-64H170.7v64z"/>`),
		g.Group(children),
	)
}
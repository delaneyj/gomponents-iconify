package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 236.7h341.3v-64H0v64zm0 106.6h341.3v-64H0v64zM0 130h341.3V66H0v64zm0 320h341.3v-64H0v64zm426.7-213.3H512v-64h-85.3v64zm0 213.3H512v-64h-85.3v64zm0-384v64H512V66h-85.3zm0 277.3H512v-64h-85.3v64z"/>`),
		g.Group(children),
	)
}
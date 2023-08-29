package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M426.7 64H384v85.3h42.7V64zM128 64H85.3v85.3H128V64zm341.3 42.7H448v64h-85.3v-64H149.3v64H64v-64H42.7C19.1 106.7 0 125.8 0 149.3v256C0 428.9 19.1 448 42.7 448h426.7c23.5 0 42.7-19.1 42.7-42.7v-256c-.1-23.5-19.2-42.6-42.8-42.6zM192 384H42.7V234.7H192V384zm149.3-21.3H234.7V320h106.7v42.7zm128-64H234.7V256h234.7v42.7z"/>`),
		g.Group(children),
	)
}
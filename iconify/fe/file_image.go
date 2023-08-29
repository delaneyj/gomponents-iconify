package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFileImage0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileImage1" fill="currentColor"><path id="feFileImage2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM15 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-7 2l3.07-3L14 16l1-1l1 1v2H8v-2Z"/></g></g>`),
		g.Group(children),
	)
}
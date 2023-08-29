package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Addstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M78 621h561c43 0 78-35 78-78V150c0-43-35-78-78-78H78c-43 0-78 35-78 78v393c0 43 35 78 78 78zm-36-78V150c0-19 17-35 36-35h561c19 0 36 16 36 35v393c0 19-17 35-36 35H78c-19 0-36-16-36-35zm94-19l114-60l114 60l-22-127l92-89l-127-20l-57-114l-56 114l-128 20l92 89zm426-177h80v45h-80v81h-46v-81h-81v-45h81v-81h46v81z"/>`),
		g.Group(children),
	)
}
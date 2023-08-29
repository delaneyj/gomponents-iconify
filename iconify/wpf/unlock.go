package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M7 0C4.79 0 2.878.917 1.687 2.406C.498 3.896 0 5.826 0 7.906V11h3V7.906c0-1.58.389-2.82 1.031-3.625C4.674 3.477 5.541 3 7 3c1.463 0 2.328.45 2.969 1.25c.64.8 1.031 2.06 1.031 3.656V9h3V7.906c0-2.092-.527-4.044-1.719-5.531C11.09.888 9.206 0 7 0zm2 10c-1.656 0-3 1.344-3 3v10c0 1.656 1.344 3 3 3h14c1.656 0 3-1.344 3-3V13c0-1.656-1.344-3-3-3H9zm7 5a2 2 0 0 1 2 2c0 .738-.404 1.372-1 1.719V21c0 .551-.449 1-1 1c-.551 0-1-.449-1-1v-2.281c-.596-.347-1-.98-1-1.719a2 2 0 0 1 2-2z"/>`),
		g.Group(children),
	)
}
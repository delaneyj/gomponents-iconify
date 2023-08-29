package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySkeleton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.853 2.159a.5.5 0 0 0-.707-.013l-1.89 1.89c-.01.009-.022.012-.03.022c-.009.008-.012.02-.02.028l-9.911 9.912A4.457 4.457 0 0 0 6.5 13a4.5 4.5 0 1 0 4.5 4.5a4.457 4.457 0 0 0-.998-2.795l6.757-6.757l2.473 2.474a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707L17.466 7.24l2.121-2.121l1.06 1.059a.498.498 0 0 0 .706 0a.5.5 0 0 0 0-.707l-1.059-1.06l1.56-1.559a.5.5 0 0 0 0-.694zM6.5 21a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7z"/>`),
		g.Group(children),
	)
}
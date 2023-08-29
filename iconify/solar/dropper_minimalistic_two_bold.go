package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropperMinimalisticTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M19 8v7.883c0 .76-.32 1.487-.88 2.001a9.024 9.024 0 0 1-5.37 2.352v1.014a.75.75 0 0 1-1.5 0v-1.014a9.024 9.024 0 0 1-5.37-2.352a2.716 2.716 0 0 1-.88-2V8c0-1.886 0-2.828.586-3.414C6.172 4 7.114 4 9 4h6c1.886 0 2.828 0 3.414.586C19 5.172 19 6.114 19 8Zm-7 5c1.105 0 2-.933 2-2.083c0-.72-.783-1.681-1.37-2.3a.862.862 0 0 0-1.26 0c-.587.619-1.37 1.58-1.37 2.3c0 1.15.895 2.083 2 2.083Zm-2.75 3a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/><path d="M13.732 3a2 2 0 0 0-3.464 0h3.464Z"/></g>`),
		g.Group(children),
	)
}
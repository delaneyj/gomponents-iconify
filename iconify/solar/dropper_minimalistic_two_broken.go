package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropperMinimalisticTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M14 16h-4"/><path d="M14 10.917c0 1.15-.895 2.083-2 2.083s-2-.933-2-2.083c0-.72.783-1.681 1.37-2.3a.862.862 0 0 1 1.26 0c.587.619 1.37 1.58 1.37 2.3Z"/><path stroke-linecap="round" d="M12 21v1m2-18a2 2 0 1 0-4 0"/><path stroke-linecap="round" d="M5 8c0-1.886 0-2.828.586-3.414C6.172 4 7.114 4 9 4h6c1.886 0 2.828 0 3.414.586C19 5.172 19 6.114 19 8v7.883c0 .76-.32 1.487-.88 2.001a9.053 9.053 0 0 1-12.24 0a2.716 2.716 0 0 1-.88-2V12"/></g>`),
		g.Group(children),
	)
}
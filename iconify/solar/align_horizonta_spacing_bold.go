package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontaSpacingBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 1.25a.75.75 0 0 1 .75.75v20a.75.75 0 0 1-1.5 0V2A.75.75 0 0 1 3 1.25Zm18 0a.75.75 0 0 1 .75.75v20a.75.75 0 0 1-1.5 0V2a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M12 4c-1.886 0-2.828 0-3.414.586C8 5.172 8 6.114 8 8v8c0 1.886 0 2.828.586 3.414C9.172 20 10.114 20 12 20c1.886 0 2.828 0 3.414-.586C16 18.828 16 17.886 16 16V8c0-1.886 0-2.828-.586-3.414C14.828 4 13.886 4 12 4Z"/></g>`),
		g.Group(children),
	)
}
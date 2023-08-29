package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontaSpacingBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M21 22.75a.75.75 0 0 1-.75-.75V2a.75.75 0 0 1 1.5 0v20a.75.75 0 0 1-.75.75Zm-18 0a.75.75 0 0 1-.75-.75V2a.75.75 0 0 1 1.5 0v20a.75.75 0 0 1-.75.75Z" clip-rule="evenodd" opacity=".5"/><path d="M12 20c1.886 0 2.828 0 3.414-.586C16 18.828 16 17.886 16 16V8c0-1.886 0-2.828-.586-3.414C14.828 4 13.886 4 12 4c-1.886 0-2.828 0-3.414.586C8 5.172 8 6.114 8 8v8c0 1.886 0 2.828.586 3.414C9.172 20 10.114 20 12 20Z"/></g>`),
		g.Group(children),
	)
}
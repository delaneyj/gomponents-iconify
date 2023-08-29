package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontaSpacingBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M3 2v20m18-12v12m0-20v4M8 15v1c0 1.886 0 2.828.586 3.414C9.172 20 10.114 20 12 20c1.886 0 2.828 0 3.414-.586C16 18.828 16 17.886 16 16V8c0-1.886 0-2.828-.586-3.414C14.828 4 13.886 4 12 4c-1.886 0-2.828 0-3.414.586C8 5.172 8 6.114 8 8v3"/>`),
		g.Group(children),
	)
}
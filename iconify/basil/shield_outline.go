package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.645 5.264a1.25 1.25 0 0 0-1.29 0l-.516.31a10.75 10.75 0 0 1-4.537 1.494l-.325.03a.25.25 0 0 0-.227.25V8.99a9.25 9.25 0 0 0 2.82 6.65l3.256 3.148a.25.25 0 0 0 .348 0l3.255-3.147a9.25 9.25 0 0 0 2.821-6.65V7.346a.25.25 0 0 0-.227-.248l-.325-.031a10.75 10.75 0 0 1-4.537-1.493l-.516-.311ZM10.58 3.979a2.75 2.75 0 0 1 2.838 0l.516.31a9.25 9.25 0 0 0 3.904 1.286l.325.03a1.75 1.75 0 0 1 1.586 1.742v1.644a10.75 10.75 0 0 1-3.278 7.73l-3.256 3.146a1.75 1.75 0 0 1-2.432 0L7.528 16.72A10.75 10.75 0 0 1 4.25 8.991V7.347a1.75 1.75 0 0 1 1.586-1.742l.325-.03a9.25 9.25 0 0 0 3.904-1.285l.516-.311Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
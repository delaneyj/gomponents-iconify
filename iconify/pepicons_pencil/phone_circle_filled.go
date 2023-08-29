package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPhoneCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="m15.874 12.546l-.91.911a31.647 31.647 0 0 1-3.535-3.537l.91-.91a2.5 2.5 0 0 0 0-3.535L10.925 4.06a2.5 2.5 0 0 0-3.536 0L4.583 6.866a.5.5 0 0 0-.114.529c2.324 6.226 6.661 10.593 13.018 13.02a.5.5 0 0 0 .531-.114l2.806-2.805a2.5 2.5 0 0 0 0-3.536l-1.414-1.414a2.5 2.5 0 0 0-3.536 0Zm4.243 2.121a1.5 1.5 0 0 1 0 2.122l-2.575 2.575c-5.821-2.306-9.811-6.32-12.023-12.02l2.577-2.576a1.5 1.5 0 0 1 2.122 0l1.414 1.414a1.5 1.5 0 0 1 0 2.121l-1.234 1.234a.5.5 0 0 0-.032.673a32.666 32.666 0 0 0 4.307 4.31a.5.5 0 0 0 .673-.031l1.235-1.236a1.5 1.5 0 0 1 2.122 0l1.414 1.414Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPhoneCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
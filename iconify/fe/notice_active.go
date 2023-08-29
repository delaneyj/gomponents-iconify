package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoticeActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feNoticeActive0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feNoticeActive1" fill="currentColor"><path id="feNoticeActive2" d="M15.085 4.853a2.501 2.501 0 1 1 2.572 3.142A5.99 5.99 0 0 1 18 10v6h1c.55 0 1 .45 1 1s-.45 1-1 1h-4v1a3 3 0 0 1-6 0v-1H5c-.55 0-1-.45-1-1s.45-1 1-1h1v-6a6.002 6.002 0 0 1 5-5.917V3a1 1 0 0 1 2 0v1.083a5.961 5.961 0 0 1 2.085.77ZM12 20a1 1 0 0 0 1-1v-1h-2v1a1 1 0 0 0 1 1Zm-4-4h8v-6a4 4 0 1 0-8 0v6Z"/></g></g>`),
		g.Group(children),
	)
}
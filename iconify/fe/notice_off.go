package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoticeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feNoticeOff0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feNoticeOff1" fill="currentColor"><path id="feNoticeOff2" d="M15 18v1a3 3 0 0 1-6 0v-1H5c-.55 0-1-.45-1-1s.45-1 1-1h1v-6a6 6 0 0 1 .257-1.743L5 7a1.414 1.414 0 1 1 2-2l11 11h1a1 1 0 0 1 0 2h-4Zm-3 2a1 1 0 0 0 1-1v-1h-2v1a1 1 0 0 0 1 1ZM8.876 4.876A5.962 5.962 0 0 1 11 4.083V3a1 1 0 0 1 2 0v1.083c2.838.476 5 2.944 5 5.917v4L8.876 4.876Z"/></g></g>`),
		g.Group(children),
	)
}
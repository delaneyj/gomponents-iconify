package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.915 14.734L23.524 24.3l-2.391-9.566l-2.392 9.566l-2.391-9.566m14.986 8.359A2.39 2.39 0 0 1 29.26 24.3h0a2.391 2.391 0 0 1-2.392-2.391v-1.555a2.391 2.391 0 0 1 2.392-2.391h0a2.391 2.391 0 0 1 2.391 2.39v.778h-4.783M15.987 23.7v9.566m0-3.946a2.391 2.391 0 0 1 2.391-2.391h0a2.391 2.391 0 0 1 2.392 2.391v3.946m9.254-3.946a2.391 2.391 0 0 1 2.391-2.391h0m-2.391 0v6.337M9.65 32.731a2.69 2.69 0 0 0 1.967.535h.537a1.583 1.583 0 0 0 1.58-1.584h0a1.583 1.583 0 0 0-1.58-1.585H11.08a1.583 1.583 0 0 1-1.58-1.584h0a1.583 1.583 0 0 1 1.58-1.584h.537a2.69 2.69 0 0 1 1.967.534m24.603 4.597a2.39 2.39 0 0 1-2.078 1.206h0a2.391 2.391 0 0 1-2.392-2.391V29.32a2.391 2.391 0 0 1 2.392-2.391h0A2.391 2.391 0 0 1 38.5 29.32v.777h-4.783m-6.051.778a2.391 2.391 0 0 1-2.391 2.39h0a2.391 2.391 0 0 1-2.391-2.39V29.32a2.391 2.391 0 0 1 2.391-2.391h0a2.391 2.391 0 0 1 2.391 2.391m0 3.946v-6.337"/>`),
		g.Group(children),
	)
}
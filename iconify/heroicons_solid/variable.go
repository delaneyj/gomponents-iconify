package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Variable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.649 3.084A1 1 0 0 1 5.163 4.4A13.95 13.95 0 0 0 4 10c0 1.993.416 3.886 1.164 5.6a1 1 0 0 1-1.832.8A15.95 15.95 0 0 1 2 10c0-2.274.475-4.44 1.332-6.4a1 1 0 0 1 1.317-.516ZM12.96 7a3 3 0 0 0-2.342 1.126l-.328.41l-.111-.279A2 2 0 0 0 8.323 7H8a1 1 0 0 0 0 2h.323l.532 1.33l-1.035 1.295a1 1 0 0 1-.781.375H7a1 1 0 1 0 0 2h.039a3 3 0 0 0 2.342-1.126l.328-.41l.111.279A2 2 0 0 0 11.677 14H12a1 1 0 1 0 0-2h-.323l-.532-1.33l1.035-1.295A1 1 0 0 1 12.961 9H13a1 1 0 1 0 0-2h-.039Zm1.874-2.6a1 1 0 0 1 1.833-.8A15.95 15.95 0 0 1 18 10c0 2.274-.475 4.44-1.332 6.4a1 1 0 1 1-1.832-.8A13.949 13.949 0 0 0 16 10c0-1.993-.416-3.886-1.165-5.6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
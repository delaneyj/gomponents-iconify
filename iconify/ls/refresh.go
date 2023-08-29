package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M536.25 399h94c0 82-31 165-92 226c-123 123-323 123-446 0s-123-323 0-446c56-56 129-86 203-91V0l241 139l-241 140v-97c-50 4-98 26-136 64c-87 86-87 226 0 312c86 87 226 87 312 0c44-43 66-102 65-159z"/>`),
		g.Group(children),
	)
}
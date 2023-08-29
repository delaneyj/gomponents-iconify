package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 8l9 7h6l9-7M15.118 3.271L3.706 6.783A1 1 0 0 0 3 7.739V23a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2V7.739a1 1 0 0 0-.706-.956L16.882 3.27a3 3 0 0 0-1.764 0Z"/>`),
		g.Group(children),
	)
}
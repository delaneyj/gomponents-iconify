package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PensiveFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><path fill="#917524" d="M53.2 25.7c-3.2-2.7-7.5-3.9-11.7-3.1c-.6.1-1.1-2-.4-2.2c4.8-.9 9.8.5 13.5 3.6c.6.5-1 2.1-1.4 1.7m-30.7-3.3c-4.2-.7-8.5.4-11.7 3.1c-.4.4-2-1.2-1.4-1.7c3.7-3.2 8.7-4.5 13.5-3.6c.7.2.2 2.4-.4 2.2"/><path fill="#664e27" d="M35.9 32.2c4.2 8 12.7 8 16.9 0c.2-.4-.3-.6-1-1c-4.2 3.3-11.1 3-14.9 0c-.6.4-1.2.6-1 1m-24.7 0c4.2 8 12.7 8 16.9 0c.2-.4-.3-.6-1-1c-4.2 3.3-11.1 3-14.9 0c-.7.4-1.2.6-1 1M38 50H26c-1.3 0-1.3-4 0-4h12c1.3 0 1.3 4 0 4"/>`),
		g.Group(children),
	)
}
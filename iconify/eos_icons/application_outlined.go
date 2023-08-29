package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="eosIconsApplicationOutlined0" fill="currentColor" d="m9 18l.7-.7l.7-.7l-1.3-1.3L7.8 14l1.3-1.3l1.3-1.3l-.7-.7L9 10l-2 2l-2 2l2 2l2 2zm4.6-1.4l.7.7l.7.7l2-2l2-2l-2-2l-2-2l-.7.7l-.7.7l1.3 1.3l1.3 1.3l-1.3 1.3l-1.3 1.3z"/></defs><path fill="currentColor" d="M1 3v2h22V3a2.001 2.001 0 0 0-2-2H3a2.001 2.001 0 0 0-2 2Zm3.007 1.008a1 1 0 1 1 .999-1a1 1 0 0 1-1 1Zm2.997-.001a1 1 0 1 1 1-1a1.002 1.002 0 0 1-1 1ZM1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm20 14H3V8h18Z"/><use href="#eosIconsApplicationOutlined0"/><use href="#eosIconsApplicationOutlined0"/>`),
		g.Group(children),
	)
}
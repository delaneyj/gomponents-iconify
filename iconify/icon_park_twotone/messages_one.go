package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessagesOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMessagesOne0"><g fill="#555" stroke="#fff" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linejoin="round" d="M34 23c0 3.862-2.703 7.157-6.5 8.433c-1.09.367-2.269.567-3.5.567c-4 0-9 2-9 2l1.132-2.446c.563-1.216.204-2.63-.57-3.723C14.572 26.436 14 24.778 14 23c0-4.97 4.477-9 10-9s10 4.03 10 9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMessagesOne0)"/>`),
		g.Group(children),
	)
}
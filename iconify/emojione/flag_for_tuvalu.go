package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTuvalu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2a5f9e" d="M12 12h21v21H12z"/><path fill="#5b97b1" d="M32 2v30H2c0 16.6 13.4 30 30 30s30-13.4 30-30S48.6 2 32 2z"/><path fill="#fff" d="M32 2c-4.7 0-9.1 1.1-13.1 3v6H11v7.9H5c-1.9 4-3 8.4-3 13.1h12V17l12 15h6v-7.5L23.6 14H32V2"/><g fill="#ed4c5c"><path d="M15.4 14L30 32h2v-4.9L21.4 14z"/><path d="M32 5H18.9c-6 2.9-11 7.9-13.9 13.9V32h6V11h21V5z"/></g><path fill="#ffe62e" d="m28.9 50.2l-.6-2.1l-.7 2.1h-2.2l1.8 1.2l-.7 2.1l1.8-1.3l1.7 1.3l-.7-2.1l1.8-1.2zm8.9-1.6l.6-2.1l-1.7 1.3l-1.7-1.3l.6 2.1l-1.7 1.2H36l.7 2.1l.7-2.1h2.1zm.6-9.3l-1.7 1.3l-1.7-1.3l.6 2l-1.7 1.3H36l.7 2.1l.7-2.1h2.1l-1.7-1.3zm6.8 7.4l.7-2l-1.8 1.3l-1.7-1.3l.7 2l-1.8 1.3h2.2l.6 2l.7-2H47zm-1.7-14.4l.6 2l.7-2H47L45.2 31l.7-2l-1.8 1.3l-1.7-1.3l.7 2l-1.8 1.3zm6.2 8.9l-.7-2l-.6 2h-2.2l1.7 1.3l-.6 2l1.7-1.2l1.8 1.2l-.7-2l1.7-1.3zm0-11l.7 2l.7-2h2.1l-1.7-1.3l.7-2.1l-1.8 1.3l-1.7-1.3l.6 2.1l-1.7 1.3zm6.3 7l-.7-2.1l-.7 2.1h-2.1l1.7 1.2l-.6 2.1l1.7-1.3l1.8 1.3l-.7-2.1l1.7-1.2zm-1.8-12.6l-.6 2l1.7-1.2l1.8 1.2l-.7-2l1.7-1.3H56l-.7-2l-.7 2h-2.1z"/>`),
		g.Group(children),
	)
}
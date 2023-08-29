package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" d="M204 0h48v48h-48z"/><path fill="#90CAF9" d="M244 45h-32V3h22l10 10z"/><path fill="#E1F5FE" d="M242.5 14H233V4.5z"/><g fill="#1976D2"><circle cx="227" cy="30" r="4"/><path d="m234 21l-5-2v11h2v-7.1l3 1.1z"/></g><path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/><g fill="#1976D2"><circle cx="23" cy="30" r="4"/><path d="m30 21l-5-2v11h2v-7.1l3 1.1z"/></g>`),
		g.Group(children),
	)
}
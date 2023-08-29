package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Direction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDirection0"><g fill="none"><rect width="36" height="36" x="6" y="6" fill="#555" stroke="#fff" stroke-width="4" rx="3"/><path fill="#fff" d="m23.293 10.565l-3.243 3.242c-.63.63-.183 1.708.707 1.708h6.486c.89 0 1.337-1.078.707-1.708l-3.243-3.242a1 1 0 0 0-1.414 0ZM10.565 24.707l3.242 3.243c.63.63 1.708.183 1.708-.707v-6.486c0-.89-1.078-1.337-1.708-.707l-3.242 3.243a1 1 0 0 0 0 1.414Zm14.142 12.728l3.243-3.242c.63-.63.183-1.708-.707-1.708h-6.486c-.89 0-1.337 1.078-.707 1.708l3.243 3.242a1 1 0 0 0 1.414 0Zm12.728-14.142l-3.242-3.243c-.63-.63-1.708-.183-1.708.707v6.486c0 .89 1.078 1.337 1.708.707l3.242-3.243a1 1 0 0 0 0-1.414Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDirection0)"/>`),
		g.Group(children),
	)
}
package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 14h13l.5-.5V9l-2.77-7.66l-.47-.34H4.27l-.47.33L1 8.74v4.76l.5.5zM14 13H2v-2.98h2.55l.74 1.25l.43.24h4.57l.44-.26l.69-1.23H14V13zm-.022-3.98H11.12l-.43.26l-.69 1.23H6.01l-.75-1.25l-.43-.24H2V9l2.62-7h6.78l2.578 7.02z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FtpServer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.6 21c1-2.4 2.3-4.2 4.1-5.7c2.2-1.9 3.8-2.5 7-3.4l-1.3-4.8c-3.4.4-6.7 1.9-9.5 4.3c-2.3 2-4 4.6-5 7.3l4.7 2.3Zm6.1 2.9c1.8-2.9 4.1-5 7.1-6l-1.2-3.7c-3.2-.2-8.8 4.7-9.5 8l3.6 1.7Zm18.8 1c3.1 3.5 2.6 8.8-1 12c-3.6 3.2-9 3.1-12.1-.4s-2.7-9 .9-12.2s9.1-2.9 12.2.6ZM22.9 35.6l12.8-11.4"/><circle cx="23.8" cy="29.5" r=".75" fill="currentColor"/><circle cx="29.5" cy="24.5" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.4 19.2l.4 3.2m-9 4.4l3.1.7m-6 6.2h0c-.7.6-.8 1.7-.1 2.4l3.5 3.9l2.5-2.2l-3.5-3.9c-.7-.8-1.7-.8-2.4-.2Zm18.7-16.6h0c-.7.6-.8 1.7-.1 2.4l3.5 3.9l2.5-2.2l-3.5-3.9c-.7-.7-1.7-.8-2.4-.2Z"/>`),
		g.Group(children),
	)
}
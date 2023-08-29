package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.03 4.83a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 0 0 1.06 1.06l2.08-2.08h11.413a4.478 4.478 0 1 0-.19-8.951a5.38 5.38 0 0 0-.437-1.834L21.03 4.83Zm-4.31 4.312L8.61 17.25h9.912a2.978 2.978 0 1 0-.77-5.854a.75.75 0 0 1-.939-.813a3.957 3.957 0 0 0-.095-1.44Z" clip-rule="evenodd"/><path fill="currentColor" d="M12.932 4.708c1.107 0 2.136.333 2.993.903a.24.24 0 0 1 .032.371l-.728.728a.261.261 0 0 1-.317.036a3.91 3.91 0 0 0-5.504 1.676a.75.75 0 0 1-.947.373a4.375 4.375 0 0 0-3.708 7.906c.152.086.19.295.067.419l-.724.723a.243.243 0 0 1-.299.038A5.875 5.875 0 0 1 8.38 7.195a5.405 5.405 0 0 1 4.552-2.487Z"/>`),
		g.Group(children),
	)
}
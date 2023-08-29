package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Custard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.84 14c-1.34 0-2.522.925-2.841 2.233l-.001.003v.001l-.003.01L6.06 20h.006l-1.427 5.76c-.02.08-.03.16-.03.24H3c-.55 0-1 .45-1 1s.45 1 1 1h1c1.28 1.28 3.02 2 4.83 2h14.34c1.81 0 3.55-.72 4.83-2h1c.55 0 1-.45 1-1s-.45-1-1-1h-1.62a1 1 0 0 0-.028-.235L25.959 20h.012l-.909-3.755A2.935 2.935 0 0 0 22.21 14H9.84Zm15.511 12H6.641l1.485-6h15.776l1.45 6Zm-2.763-9.92a.908.908 0 0 1 .517.624l.003.01l.31 1.286H8.621l.315-1.271l.004-.017l.001-.005A.937.937 0 0 1 9.84 16h12.37c.134 0 .262.029.378.08Z"/>`),
		g.Group(children),
	)
}
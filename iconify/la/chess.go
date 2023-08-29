package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 4v2H4.8l1.274 6H6c-.6 0-1 .4-1 1s.4 1 1 1h.11l-1.018 7H5c-.6 0-1 .4-1 1c0 .32.12.576.316.752L3 24.699V27h12.1v-2.3l-1.356-2.007A1 1 0 0 0 14 22c0-.6-.4-1-1-1h-.11l-1.081-7H12c.6 0 1-.4 1-1s-.4-1-1-1h-.05l1.35-6H10V4H8zm-.8 4h3.6l-.9 4H8.1l-.9-4zM19 9v6.4l.9.9l-.771 4.7H19c-.6 0-1 .4-1 1c0 .32.12.576.316.752L17 24.699V27h12.1v-2.3l-1.356-2.007A1 1 0 0 0 28 22c0-.6-.4-1-1-1h-.129l-.771-4.7l.9-.9V9h-2v2h-1V9h-2v2h-1V9h-2zm2 4h4v1.6l-1.1 1.1l.9 5.3h-3.6l.9-5.3l-1.1-1.1V13zM8.2 14h1.7l1 7H7.2l1-7zm-1.7 9h5l1.3 2H5.2l1.3-2zm14.1 0h4.9l1.3 2h-7.6l1.4-2z"/>`),
		g.Group(children),
	)
}
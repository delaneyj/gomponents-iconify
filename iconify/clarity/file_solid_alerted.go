package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M22.2 15.4c-2 0-3.7-1.6-3.7-3.6c0-.7.2-1.3.5-1.9l3.2-5.5l-.3-.4H7.8C6.8 4 6 4.9 6 5.9v24.2c0 1 .8 1.9 1.8 1.9h20.3c1 0 1.8-.9 1.8-1.9V15.4h-7.7z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M26.9 1.1L21.1 11c-.4.6-.2 1.4.3 1.8c.2.2.5.2.8.2h11.5c.7 0 1.3-.5 1.3-1.2c0-.3-.1-.5-.2-.8l-5.7-9.9C28.7.5 28 .3 27.3.6c-.2.2-.3.4-.4.5z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}
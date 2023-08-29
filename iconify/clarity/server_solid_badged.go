package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M24 10.3v1.2H12V10h11.8c-.5-.7-.8-1.5-1-2.4H12V6h10.5v-.1c0-1.4.4-2.7 1.1-3.9H9.5C8.7 2 8 2.7 8 3.5V34h20V13.1c-1.6-.5-3-1.4-4-2.8zm-6 20.2c-1.5 0-2.8-1.2-2.8-2.8S16.5 25 18 25s2.8 1.2 2.8 2.8s-1.3 2.7-2.8 2.7zm5-7.9H13V21h10v1.6z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="18" cy="27.8" r="1.2" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><circle cx="30" cy="5.9" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}
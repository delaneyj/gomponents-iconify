package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instlife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="4.4" d="M6.115 32.255c-1.123-.39-2.61-1.637-2.61-2.19c0-.134.806-.222 2.042-.222c4.4 0 6.927-1.498 9.747-5.78c2.079-3.157 4.204-4.313 8.042-4.375l2.34-.038l1.113-1.113c2.234-2.234 4.874-3.207 8.658-3.19c2.767.013 4.189.388 6.079 1.608c2.628 1.695 3.588 4.239 2.598 6.88c-1.281 3.417-5.505 5.79-10.307 5.79c-3.907 0-6.873-1.382-8.49-3.958l-.498-.793l-1.155.208c-1.7.306-4.252 1.205-5.576 1.966c-.639.367-1.91 1.369-2.825 2.227c-1.898 1.78-3.113 2.554-4.799 3.056c-1.539.458-2.885.435-4.358-.077h0Z"/>`),
		g.Group(children),
	)
}
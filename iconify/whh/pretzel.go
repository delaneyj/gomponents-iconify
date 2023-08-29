package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pretzel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 768q-103 0-198.5-35t-164-92.5t-109-134T0 349q0-62 17.5-123.5t48-112t77.5-82T244 0q113 0 268 116Q667 0 780 0q54 0 101 31.5t77.5 82t48 112T1024 349q0 81-40.5 157.5t-109 134t-164 92.5T512 768zm202-172q-26-91-79.5-178.5T511 262q-69 67-122 155t-79 179q94 44 202 44t202-44zM128 361q0 82 67 153q32-91 87.5-175.5T410 191q-106-63-154-63q-34 0-59.5 20.5t-40 55.5t-21.5 74.5t-7 82.5zm640-233q-48 0-154 63q72 63 127.5 147.5T829 514q67-71 67-153q0-58-13-109t-43-87.5t-72-36.5z"/>`),
		g.Group(children),
	)
}
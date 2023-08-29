package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.7 2.3L2.3 3.7l6.821 6.82c-.006.027-.015.052-.021.08l1.9 1.9v-.102L15.602 17H15.5l2.2 2.2c.05.01.096.032.146.044l5.814 5.815c.01.048.03.092.04.14L25.5 27h.102l2.699 2.7l1.398-1.4l-4.105-4.105c-.844-2.88-2.946-5.25-5.694-6.394C21.8 16.5 23 14.4 23 12c0-3.9-3.1-7-7-7c-2.609 0-4.853 1.42-6.078 3.523L3.699 2.301zM16 7c2.8 0 5 2.2 5 5c0 2.087-1.224 3.838-3.006 4.596l-6.59-6.59C12.162 8.224 13.913 7 16 7zm-6.9 6.3c.4 1.9 1.4 3.5 3 4.5C8.5 19.3 6 22.9 6 27h2c0-4.1 3-7.4 6.9-7.9l-5.8-5.8z"/>`),
		g.Group(children),
	)
}
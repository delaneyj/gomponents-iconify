package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M3.5 21.501v17c0 2.49.561 2.999 3 2.999h28c2.609 0 3-.471 3-2.999v-17c0-2.459-.46-3.001-3-3.001h-28c-2.41 0-3 .671-3 3.001zm20.18 7c0 1.41-.91 2.599-2.17 3.019l1.99 4.98h-6l1.99-4.98a3.176 3.176 0 0 1-2.17-3.019c0-1.76 1.42-3.181 3.18-3.181s3.18 1.421 3.18 3.181zm-17.18-12l5.41-.029c.26-4.78 2.439-9.631 8.64-9.631c4.56 0 6.851 2.74 8.24 5.99l4.771-3.359c-1.939-4.631-6.431-7.91-12.811-7.971c-9.27-.09-14.25 6.641-14.25 15z"/>`),
		g.Group(children),
	)
}
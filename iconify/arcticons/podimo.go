package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podimo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.946 35.744V6.038c0-.85.689-1.538 1.538-1.538h0c8.597 0 15.566 6.969 15.566 15.566v1.65c0 8.597-6.97 15.566-15.566 15.566h0c-.85 0-1.538-.688-1.538-1.538ZM7.95 27.217V5.618c0-.617.501-1.118 1.119-1.118h0c6.25 0 11.317 5.067 11.317 11.317v1.201c0 6.25-5.067 11.317-11.317 11.317h0c-.618 0-1.119-.5-1.119-1.118Z"/><circle cx="14.168" cy="37.282" r="6.218" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
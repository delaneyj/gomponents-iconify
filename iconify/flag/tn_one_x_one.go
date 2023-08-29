package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TnOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#e70013" d="M0 0h512v512H0z"/><path fill="#fff" d="M256 135a1 1 0 0 0-1 240a1 1 0 0 0 0-241zm72 174a90 90 0 1 1 0-107a72 72 0 1 0 0 107zm-4.7-21.7l-37.4-12.1l-23.1 31.8v-39.3l-37.3-12.2l37.3-12.2v-39.4l23.1 31.9l37.4-12.1l-23.1 31.8z"/>`),
		g.Group(children),
	)
}
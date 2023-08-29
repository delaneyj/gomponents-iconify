package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PkOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPk1x10"><path fill-opacity=".7" d="M0 0h512v512H0z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagPk1x10)"><path fill="#0c590b" d="M-95 0h768v512H-95z"/><path fill="#fff" d="M-95 0H97.5v512H-95z"/><g fill="#fff"><path d="m403.7 225.4l-31.2-6.6l-16.4 27.3l-3.4-31.6l-31-7.2l29-13l-2.7-31.7l21.4 23.6l29.3-12.4l-15.9 27.6l21 24z"/><path d="M415.4 306a121.2 121.2 0 0 1-161.3 59.4a122.1 122.1 0 0 1-59.5-162.1A118.6 118.6 0 0 1 266 139a156.2 156.2 0 0 0-11.8 10.9A112.3 112.3 0 0 0 415.5 306z"/></g></g>`),
		g.Group(children),
	)
}
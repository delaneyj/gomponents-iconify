package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DjOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagDj1x10"><path fill-opacity=".7" d="M55.4 0H764v708.7H55.4z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagDj1x10)" transform="translate(-40) scale(.722)"><path fill="#0c0" d="M0 0h1063v708.7H0z"/><path fill="#69f" d="M0 0h1063v354.3H0z"/><path fill="#fffefe" d="m0 0l529.7 353.9L0 707.3V0z"/><path fill="red" d="m221.2 404.3l-42.7-30.8l-42.4 31l15.8-50.3l-42.4-31.2l52.4-.4l16.3-50.2l16.6 50l52.4.2l-42.1 31.4l16 50.3z"/></g>`),
		g.Group(children),
	)
}
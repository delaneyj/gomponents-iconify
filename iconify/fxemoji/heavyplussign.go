package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heavyplussign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#2B3B47" d="M468.433 210.554H301.446V43.567c0-17.164-13.914-31.079-31.079-31.079h-28.735c-17.164 0-31.079 13.914-31.079 31.079v166.987H43.567c-17.164 0-31.079 13.914-31.079 31.079v28.735c0 17.164 13.914 31.079 31.079 31.079h166.987v166.987c0 17.164 13.914 31.079 31.079 31.079h28.735c17.164 0 31.079-13.914 31.079-31.079V301.446h166.987c17.164 0 31.079-13.914 31.079-31.079v-28.735c-.001-17.164-13.916-31.078-31.08-31.078z"/>`),
		g.Group(children),
	)
}
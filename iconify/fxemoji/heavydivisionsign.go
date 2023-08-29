package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heavydivisionsign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#2B3B47" d="M12.488 270.368v-28.735c0-17.164 13.914-31.079 31.079-31.079h424.866c17.164 0 31.079 13.914 31.079 31.079v28.735c0 17.164-13.914 31.079-31.079 31.079H43.567c-17.164-.001-31.079-13.915-31.079-31.079z"/><circle cx="256" cy="90.258" r="56.247" fill="#2B3B47"/><circle cx="256" cy="421.742" r="56.247" fill="#2B3B47"/>`),
		g.Group(children),
	)
}
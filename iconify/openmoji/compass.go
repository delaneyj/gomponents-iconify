package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M33.2 33.2L48 24l-9.2 14.8"/><circle cx="36" cy="36" r="24" fill="#fcea2b"/><path fill="#f1b31c" d="M53 19a24.042 24.042 0 0 1-17 41a24.302 24.302 0 0 1-17-7"/><path fill="#fff" d="M33.2 33.2L24 48l14.8-9.2"/><path fill="#ea5a47" d="M33.2 33.2L48 24l-9.2 14.8"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="24"/><path d="M33.2 33.2L24 48l14.8-9.2"/><path d="M33.2 33.2L48 24l-9.2 14.8M36 21v-5m0 40v-5m15-15h5m-40 0h5"/></g>`),
		g.Group(children),
	)
}
package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FCEA2B" d="m55.532 51.262l-19.02 1l-19.021-1l-.903-18.886H56z"/><path fill="#3F3F3F" d="M55.355 32.376L56.5 7.083h-40l1.145 25.293zM54.5 51.262s-11.29 1.77-17.5 1.77s-18.5-1.77-18.5-1.77c-7.313-1.159-14-2.076-14-.04c0 3.56 14.327 6.447 32 6.447s32-2.886 32-6.447c0-2.036-6.687-1.119-14 .04z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m55.355 32.376l-.673 14.876L56.5 7.083h-40l1.818 40.169l-.673-14.876zM54.5 51.262s-11.29 1.77-17.5 1.77s-18.5-1.77-18.5-1.77c-7.313-1.159-14-2.076-14-.04c0 3.56 14.327 6.447 32 6.447s32-2.886 32-6.447c0-2.036-6.687-1.119-14 .04z"/>`),
		g.Group(children),
	)
}
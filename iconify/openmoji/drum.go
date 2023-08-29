package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#EA5A47" d="M58.757 59.067V37.22H13.019v21.848h45.738z"/><path fill="#a57939" d="M23.9 30.276c.29.385.715.635 1.194.702c.483.06.955-.058 1.34-.348v-.002a1.81 1.81 0 0 0-1.09-3.254c-.38 0-.763.118-1.088.364a1.815 1.815 0 0 0-.355 2.538zm22.623.264a2.011 2.011 0 0 0-.329-2.82a2.007 2.007 0 1 0-2.49 3.148c.423.331.95.49 1.477.42a2.003 2.003 0 0 0 1.343-.749z"/><path fill="#D22F27" d="m47.681 54.448l-.057 4.62h11.133v-21.85H47.894l-.08 6.513l.08-.003a1 1 0 0 1 1 1v8.744a1 1 0 0 1-1 1l-.213-.024z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M35.815 44.728v8.744m-12.079-8.744v8.744m24.158-8.744v8.744"/><circle cx="25.34" cy="29.179" r="2" transform="rotate(-37.022 25.341 29.181)"/><path d="m38.775 19.046l11.518-8.687M26.937 27.975l5.041-3.802"/><circle cx="44.944" cy="29.234" r="2" transform="rotate(-51.655 44.946 29.234)"/><path d="m21.337 10.562l22.038 17.431m-31.357 8.226h47.739v23.849H12.018z"/></g>`),
		g.Group(children),
	)
}
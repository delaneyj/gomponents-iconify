package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomansClothes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ff99b6" d="M1.7 24.1L.6 11.3s8.7-3.6 11.7-4.1c4.8-.8 14-2 19.7-2s14.8 1.2 19.7 2c3.1.5 11.7 4.1 11.7 4.1l-1.1 12.8l-12.1.9c-.2 13.7 1.9 30.5 2.6 36.1c0 0-8.9.4-20.8.5c-11.9-.1-20.8-.5-20.8-.5c.7-5.6 2.8-22.4 2.6-36.1l-12.1-.9"/><g fill="#ff78a5"><path d="M32 12.4h1.2v49.2H32z"/><path d="M26.3 5.6h11.5v7H26.3z"/></g><path fill="#ff99b6" d="M21.8 0h20.4v6.6H21.8z"/><g fill="#ffd2df"><path d="m42.2 0l3.9 6.2l-7.4 11.5l-6.7-5.1z"/><path d="m21.8 0l-3.9 6.2l7.4 11.5l6.7-5.1z"/></g><circle cx="28.9" cy="20.4" r="2.6" fill="#ff78a5"/><ellipse cx="28.9" cy="20.1" fill="#fff3f7" rx="2.5" ry="2.3"/><circle cx="28.9" cy="28.9" r="2.6" fill="#ff78a5"/><ellipse cx="28.9" cy="28.6" fill="#fff3f7" rx="2.5" ry="2.3"/><circle cx="28.9" cy="37.4" r="2.6" fill="#ff78a5"/><ellipse cx="28.9" cy="37.1" fill="#fff3f7" rx="2.5" ry="2.3"/><circle cx="28.9" cy="45.9" r="2.6" fill="#ff78a5"/><ellipse cx="28.9" cy="45.6" fill="#fff3f7" rx="2.5" ry="2.3"/><circle cx="28.9" cy="54.4" r="2.6" fill="#ff78a5"/><ellipse cx="28.9" cy="54.1" fill="#fff3f7" rx="2.5" ry="2.3"/><path fill="#ffd2df" d="M64 11.1L60.7 10l-1.2 14.6l3.5-.3zm-64 0L3.3 10l1.2 14.6l-3.5-.3zm32 48c10.7-.2 19.2-1.6 20.6-2.5c.1 1.9.1 3.5.2 4.8c0 0-8.9 2.5-20.8 2.7c-11.9-.2-20.8-2.7-20.8-2.7c.1-1.2.2-2.9.2-4.8c1.4.8 9.9 2.3 20.6 2.5"/><path fill="#ff99b6" d="M32 59.1h1.2V64H32z"/>`),
		g.Group(children),
	)
}
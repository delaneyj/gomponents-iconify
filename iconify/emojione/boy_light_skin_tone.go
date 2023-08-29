package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyLightSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#594640" d="M32 2c23 0 28 15.3 28 24c0 7.7-1 11-1 11H5s-1-3.3-1-11C4 17.3 9 2 32 2z"/><path fill="#ffe1bd" d="M57.2 33s-.8-3.3-1.2-10.2c-.5-7.5-10.9-1.5-24-1.5s-23.5-6-24 1.5C7.6 29.7 6.8 33 6.8 33c-6.5 0-6.5 10.6.4 10.6C7.2 55 21.5 62 32 62s24.8-7 24.8-18.4c6.9 0 6.9-10.6.4-10.6"/><circle cx="44.5" cy="35.5" r="6.5" fill="#fff"/><circle cx="44.5" cy="35.5" r="4.5" fill="#664e27"/><circle cx="44.5" cy="35.5" r="1.5" fill="#231f20"/><circle cx="19.5" cy="35.5" r="6.5" fill="#fff"/><circle cx="19.5" cy="35.5" r="4.5" fill="#664e27"/><circle cx="19.5" cy="35.5" r="1.5" fill="#231f20"/><path fill="#664e27" d="M40.1 48.1c-5.2 3.6-11 3.6-16.2 0c-.6-.4-1.2.3-.8 1c1.6 2.6 4.8 4.9 8.9 4.9c4.1 0 7.3-2.3 8.9-4.9c.4-.7-.2-1.4-.8-1"/>`),
		g.Group(children),
	)
}
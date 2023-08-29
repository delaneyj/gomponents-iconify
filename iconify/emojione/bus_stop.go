package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#83bf4f" d="M26 26h4v38h-4z"/><path fill="#42ade2" d="M31 30h19v6H31z"/><path fill="#fc97b2" d="M31 36h19v6H31z"/><path fill="#42ade2" d="M14 30h11v6H14z"/><path fill="#fc97b2" d="M14 36h11v6H14z"/><path fill="#42ade2" d="M31 42h19v6H31z"/><path fill="#699635" d="M26 48h4v2h-4zm-1-13h6v2h-6zm0 6h6v2h-6z"/><path fill="#42ade2" d="M28 0c-7.7 0-14 6.3-14 14c0 3.9 1.6 7.4 4.1 9.9l9.9 1l9.9-1c2.5-2.5 4.1-6 4.1-9.9c0-7.7-6.3-14-14-14"/><path fill="#ff717f" d="m37.9 23.9l-2.8-2.8C31.2 25 24.9 25 21 21.1l-2.8 2.8c2.2 2.2 5 3.5 7.9 3.9V30h4v-2.2c2.8-.4 5.6-1.7 7.8-3.9"/><path fill="#fff" d="M34 15c0 1.1-.9 2-2 2h-8c-1.1 0-2-.9-2-2V7c0-1.1.9-2 2-2h8c1.1 0 2 .9 2 2v8"/><g fill="#42ade2"><path d="M32 10c0 .6-.4 1-1 1h-6c-.5 0-1-.4-1-1V8c0-.6.5-1 1-1h6c.6 0 1 .4 1 1v2"/><circle cx="25" cy="14" r="1"/><circle cx="31" cy="14" r="1"/></g><path fill="#fff" d="M31 16v3c0 .5.5 1 1 1c.6 0 1-.5 1-1v-3h-2m-8 0v3c0 .5.5 1 1 1c.6 0 1-.5 1-1v-3h-2"/><path fill="#699635" d="M25 42h6v6h-6zm0-6h6v6h-6zm0-6h6v6h-6z"/><path fill="#83bf4f" d="M27 30h2v5h-2zm0 7h2v4h-2zm0 6h2v5h-2z"/><path fill="#fff" d="M15.9 31.6h1.6v3h-1.6zm3.6 0h3.3v3h-3.3zm-3.6 5.9h3.3v3h-3.3zm5.3 0h1.7v3h-1.7zm21.4-6h6.2v1h-6.2zm-7.4 0h6.2v1h-6.2zm3.4 2h10.1v1H38.6zm-1.2 10h11.3v1H37.4zm7.6 2h3.7v1H45zm-11.7 0h10.6v1H33.3zm8.2-8h7.2v1h-7.2zm-8.6 0h7.2v1h-7.2zm12.1 2h3.8v1H45zm-4.9 0h3.8v1h-3.8zm-4.6 0h3.8v1h-3.8z"/>`),
		g.Group(children),
	)
}
package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagKenya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path d="M5 17h62v13H5z"/><path fill="#d22f27" stroke="#fff" stroke-miterlimit="10" stroke-width="2" d="M5 30.5h62v11H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m29.25 47.691l13.5-23.382m-13.5 0l13.5 23.382"/><path fill="#d22f27" d="M36.231 25.057a.496.496 0 0 0-.462 0A12.299 12.299 0 0 0 29.5 36a12.299 12.299 0 0 0 6.267 10.943a.5.5 0 0 0 .231.057h.004a.5.5 0 0 0 .23-.057A12.299 12.299 0 0 0 42.5 36a12.299 12.299 0 0 0-6.269-10.943Z"/><path d="M41.703 31.501A6.288 6.288 0 0 0 40 36a6.288 6.288 0 0 0 1.704 4.5a13.104 13.104 0 0 0 0-8.999Zm-11.407 8.998A6.288 6.288 0 0 0 32 36a6.288 6.288 0 0 0-1.703-4.499a13.103 13.103 0 0 0 0 8.998Z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M36 27v18"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}
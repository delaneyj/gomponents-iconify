package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miniangry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M959.5 897q-26.5 0-45-19T896 833q0-49-51-93t-140.5-71.5T512 641q-102 0-192 28.5T179 742t-51 91q0 26-19 45t-45.5 19t-45-19T0 833q0-94 69-167.5t185.5-113T512 513t257.5 39.5t185.5 113t69 167.5q0 26-19 45t-45.5 19zm-77-591q-13.5 15-33 15T816 306l-145-97q-11-4-18-12q-14-15-14-36.5t14-36.5q7-7 18-11l145-98q14-15 33.5-15t33 15T896 51.5T882 88l-108 73l108 72q14 15 14 36.5T882.5 306zM353 209l-145 97q-14 15-33.5 15t-33-15t-13.5-36.5t14-36.5l108-72l-108-73q-14-15-14-36.5T141.5 15t33-15T208 15l145 98q11 4 18 11q14 15 14 36.5T371 197q-7 8-18 12z"/>`),
		g.Group(children),
	)
}
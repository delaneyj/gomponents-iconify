package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 3.25a4.25 4.25 0 1 0 0 8.5a4.25 4.25 0 0 0 0-8.5ZM7.25 7.5a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M3.75 17A2.25 2.25 0 0 1 6 14.75h.34c.027 0 .053.004.078.012l.866.283a8.75 8.75 0 0 0 3.071.425a.303.303 0 0 0 .28-.244c.074-.34.176-.669.305-.985a.22.22 0 0 0-.223-.3a7.251 7.251 0 0 1-2.967-.322l-.866-.283a1.752 1.752 0 0 0-.543-.086H6A3.75 3.75 0 0 0 2.25 17v1.188c0 .754.546 1.396 1.29 1.517c2.595.424 5.221.59 7.84.5c.163-.005.251-.194.16-.328a5.998 5.998 0 0 1-.508-.903a.42.42 0 0 0-.387-.25a38.6 38.6 0 0 1-6.864-.5a.037.037 0 0 1-.031-.036V17Z"/><path fill="currentColor" fill-rule="evenodd" d="M12 16.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 21a4.5 4.5 0 1 0-4.5-4.5Zm4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 19.5Zm-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
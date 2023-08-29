package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firecracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiFirecracker0" d="m62 23l5.657-1.414M55 14.343L56.414 20M60 20l3-5"/></defs><use href="#openmojiFirecracker0" fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2"/><rect width="40" height="13" x="3.516" y="31.28" fill="#d22f27" rx="1" transform="rotate(-45 23.516 37.78)"/><use href="#openmojiFirecracker0" fill="none" stroke="#fcea2b" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><rect width="40" height="13" x="3.516" y="31.28" rx="1" transform="rotate(-45 23.516 37.78)"/><path stroke-linecap="round" d="M58 23.28s-2.757 6.829-7 4s.472-8.764-4-11c-4-2-9.342 7.357-9.342 7.357"/></g>`),
		g.Group(children),
	)
}
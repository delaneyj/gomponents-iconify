package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiMagnet0" d="M20.5 31.5h6v5h-6zm23 0h6v5h-6z"/></defs><path fill="#ea5a47" d="M43.5 36.5V47a8.5 8.5 0 0 1-17 0V36.5h-6V47a14.5 14.5 0 0 0 29 0V36.5Z"/><use href="#openmojiMagnet0" fill="#d0cfce"/><path fill="#fcea2b" d="M44.882 19.775a1.283 1.283 0 0 0-1.26-.907h-4.925l3.461-7.75a1 1 0 0 0-1.645-1.089L29.42 21.958a1.327 1.327 0 0 0-.278 1.451a1.37 1.37 0 0 0 1.266.85h4.88l-4.996 8.05a1 1 0 0 0 .213 1.299c.227.188.38.315.572.315c.4 0 .966-.557 2.716-2.279L44.21 21.371a1.452 1.452 0 0 0 .672-1.596Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M43.5 36.5V47a8.5 8.5 0 0 1-17 0V36.5h-6V47a14.5 14.5 0 0 0 29 0V36.5Z"/><use href="#openmojiMagnet0" stroke-linecap="round" stroke-linejoin="round"/><path stroke-miterlimit="10" d="M44.882 19.775a1.283 1.283 0 0 0-1.26-.907h-4.925l3.461-7.75a1 1 0 0 0-1.645-1.089L29.42 21.958a1.327 1.327 0 0 0-.278 1.451a1.37 1.37 0 0 0 1.266.85h4.88l-4.996 8.05a1 1 0 0 0 .213 1.299c.227.188.38.315.572.315c.4 0 .966-.557 2.716-2.279L44.21 21.371a1.452 1.452 0 0 0 .672-1.596Z"/></g>`),
		g.Group(children),
	)
}
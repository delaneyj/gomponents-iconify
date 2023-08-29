package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chestnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiChestnut0" stroke-miterlimit="10" d="M46.473 27.544s4.606 13.533-1.184 25.301"/></defs><path fill="#a57939" d="m28 28.614l18-1s5 25-1 25l-9 5l-7.79-4.769C24 40.614 28 28.614 28 28.614Z"/><path fill="#f1b31c" d="M17 23.614s.422-2.317 4-5c4-3 16-5 16-5l-14 12Z"/><path fill="#6a462f" d="m27 27.614l-11.25-4s-12.75 19 11.25 28Zm20 0l11.25-4s10.5 11.429-11.25 28l1-13Z"/><path fill="#f1b31c" d="m16.927 23.498l.084.194C34.508 34.184 57 23.614 57 23.614a23.843 23.843 0 0 0-19.356-10.108"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-miterlimit="10" d="m16.927 23.498l.084.194C34.508 34.184 57 23.614 57 23.614a23.843 23.843 0 0 0-19.356-10.108"/><path stroke-miterlimit="10" d="M16 23.614c17.497 10.492 41 0 41 0"/><use href="#openmojiChestnut0" stroke-miterlimit="10"/><use href="#openmojiChestnut0" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.977 17.807c-6.157 3.983-10.244 10.39-9.866 17.742c.389 7.565 6.78 14.68 13.547 16.362c4.101 1.018 8.704 4.838 10.017 5.983a.625.625 0 0 0 .773.04a69.433 69.433 0 0 1 8.549-4.984c8.574-4.202 13.503-9.836 14.977-17.4c2.29-11.762-10.085-22.043-23.302-22.043a27.39 27.39 0 0 0-11.383 2.485"/><path stroke-miterlimit="10" d="M27.026 27.544a113.756 113.756 0 0 0 1.184 25.301"/></g>`),
		g.Group(children),
	)
}
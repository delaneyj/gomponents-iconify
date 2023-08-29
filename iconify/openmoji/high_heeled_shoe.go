package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighHeeledShoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiHighHeeledShoe0" d="M25.501 34.561c2.916 2.372 5.458 5.385 6.764 9.15c4.333 12.493 1.03 12.728 4.333 17.345c3.98 5.56 40.725 5.177 23.375-7.08c0 0-12.254 2.98-16.94-1.355c0 0-5.427-17.955-26.14-38.379c0 0-4.5 2.328-6.243 8.046"/></defs><path fill="#D22F27" d="M25.501 34.561c2.916 2.372 5.458 5.385 6.764 9.15c4.333 12.493 1.03 12.728 4.333 17.345c3.98 5.56 40.725 5.177 23.375-7.08c0 0-12.254 2.98-16.94-1.355c0 0-5.427-17.955-26.14-38.379c0 0-4.5 2.328-6.243 8.046c0 0 .56 2.002.56 4.579s14.291 7.694 14.291 7.694z"/><path fill="#D22F27" d="M25.108 35.453c2.915 2.372 5.851 4.493 7.157 8.257c4.333 12.494 1.03 12.73 4.333 17.346c3.98 5.56 40.725 5.177 23.375-7.08c0 0-12.254 2.98-16.94-1.355c0 0-5.427-17.955-26.14-38.379c0 0-5.504 2.134-7.246 7.851c0 0-.644 2.13-.11 3.364s15.57 9.996 15.57 9.996z"/><path fill="#3F3F3F" d="m19.962 33.086l-2 27.872h-3.84V32.086l-2.912-5.982s6.65 3.811 10.526 6.911l-1.774.07z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m19.962 33.086l-2 27.872h-3.84V32.086l-2.912-5.982s6.65 3.811 10.526 6.911l-1.774.07z"/><use href="#openmojiHighHeeledShoe0"/><use href="#openmojiHighHeeledShoe0"/></g>`),
		g.Group(children),
	)
}